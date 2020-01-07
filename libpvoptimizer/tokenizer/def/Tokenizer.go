package def

import (
	"fmt"
	"sync"

	tokenizer ".."
	libpvoptimizer "../.."
	"../../errors"
	"../../sourcereader"
)

const (
	initialMaxStreams   = 32
	maxStreamsIncrement = 8
)

type stream struct {
	data     []byte
	line     int
	char     int
	filename string
}

// Tokenizer is the default implementation of ITokenizer
type Tokenizer struct {
	handler      errors.IErrorHandler
	sourcereader libpvoptimizer.ISourceReader
	lexer        libpvoptimizer.ILexer
	streams      []stream
}

// Init initializes the layer and is called from the pipeline layer
func (tk *Tokenizer) Init(sourcereader libpvoptimizer.ISourceReader, lexer libpvoptimizer.ILexer, e errors.IErrorHandler, wg *sync.WaitGroup) {
	tk.handler = e
	tk.sourcereader = sourcereader
	tk.lexer = lexer
	tk.streams = make([]stream, initialMaxStreams)
}

// Stream accepts a buffer and tokenizes the contents, then streams those
// tokens to the lexer.  The id is used to identify different files that are
// being tokenized at the same time (for includes).
func (tk *Tokenizer) Stream(data []byte, l int, id int) {
	s := tk.streams[id]
	buf := append(s.data, data[0:l]...)
	wasSlash := false
	wasArrow := false
	commentLine := false
	for ; len(buf) > 0; s.char++ {
		if buf[0] == '\n' {
			commentLine = false
			s.line++
			s.char = 1
			buf = buf[1:]
		} else if commentLine {
			buf = buf[1:]
		} else {
			if buf[0] == '/' {
				if wasSlash {
					commentLine = true
					wasSlash = false
				} else {
					wasSlash = true
				}
				buf = buf[1:]
			} else {
				if wasSlash {
					tk.lexer.Stream(tokenizer.Token{
						Text:     "/",
						LineNo:   s.line,
						CharNo:   s.char - 1,
						FileName: s.filename,
					}, id)
					wasSlash = false
				}
				if wasArrow {
					wasArrow = false
					if buf[0] == '-' {
						tk.lexer.Stream(tokenizer.Token{
							Text:     "<-",
							LineNo:   s.line,
							CharNo:   s.char - 1,
							FileName: s.filename,
						}, id)
						buf = buf[1:]
						continue
					} else {
						tk.lexer.Stream(tokenizer.Token{
							Text:     "<",
							LineNo:   s.line,
							CharNo:   s.char - 1,
							FileName: s.filename,
						}, id)
					}
				}
				if buf[0] == ' ' || buf[0] == '\t' || buf[0] == '\r' {
					buf = buf[1:]
				} else if (buf[0] >= 'a' && buf[0] <= 'z') || (buf[0] >= 'A' && buf[0] <= 'Z') || buf[0] == '_' {
					b := buf
					for len(buf) > 0 {
						if (buf[0] >= 'a' && buf[0] <= 'z') || (buf[0] >= 'A' && buf[0] <= 'Z') || buf[0] == '_' || (buf[0] >= '0' && buf[0] <= '9') {
							buf = buf[1:]
						} else {
							r := len(b) - len(buf)
							tk.lexer.Stream(tokenizer.Token{
								Text:     string(b[0:r]),
								LineNo:   s.line,
								CharNo:   s.char,
								FileName: s.filename,
							}, id)
							s.char += r - 1
							break
						}
					}
					if len(buf) == 0 {
						buf = b
						break
					}
				} else if buf[0] == '.' || buf[0] == '-' || (buf[0] >= '0' && buf[0] <= '9') {
					b := buf
					for len(buf) > 0 {
						if buf[0] == '.' || buf[0] == '-' || (buf[0] >= '0' && buf[0] <= '9') {
							buf = buf[1:]
						} else {
							r := len(b) - len(buf)
							tk.lexer.Stream(tokenizer.Token{
								Text:     string(b[0:r]),
								LineNo:   s.line,
								CharNo:   s.char,
								FileName: s.filename,
							}, id)
							s.char += r - 1
							break
						}
					}
					if len(buf) == 0 {
						buf = b
						break
					}
				} else if buf[0] == '<' {
					wasArrow = true
					buf = buf[1:]
				} else {
					tk.lexer.Stream(tokenizer.Token{
						Text:     fmt.Sprintf("%c", buf[0]),
						LineNo:   s.line,
						CharNo:   s.char,
						FileName: s.filename,
					}, id)
					buf = buf[1:]
				}
			}
		}
	}
	if wasSlash {
		s.data = []byte{'/'}
	} else {
		s.data = buf
	}
	tk.streams[id] = s
}

// BeginStream sets up buffers for the streaming process.  This method
// should be called from the source reading layer, and the filename should
// be as complete as possible since it is used for printing errors.
func (tk *Tokenizer) BeginStream(filename string, id int) {
	if id >= len(tk.streams) {
		tk.streams = append(tk.streams, make([]stream, maxStreamsIncrement)...)
	}
	tk.streams[id] = stream{
		data:     make([]byte, 0),
		line:     1,
		char:     1,
		filename: filename,
	}
}

// EndStream is called by the source reading layer once the entire file has
// been streamed into the tokenizer.
func (tk *Tokenizer) EndStream(id int) {
	tk.lexer.EndStream(id)
}

// ReadFile starts the reading of a different source file.  This is used
// when reading an included file.  The id is used in the lexer to
// reconstruct all the file parts in order.  This call is proxied down to
// the source reader layer.
func (tk *Tokenizer) ReadFile(filename string, id int, sourceType sourcereader.SourceType) {
	tk.sourcereader.ReadFile(filename, id, sourceType)
}
