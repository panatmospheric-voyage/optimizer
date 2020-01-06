package def

import (
	libpvoptimizer "../.."
	"../../errors"
	"../../sourcereader"
)

// Tokenizer is the default implementation of ITokenizer
type Tokenizer struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (tk Tokenizer) Init(sourcereader libpvoptimizer.ISourceReader, lexer libpvoptimizer.ILexer, e errors.IErrorHandler) {
	tk.handler = e
	errors.NoImpl(tk.handler, "Tokenizer.Init")
}

// Stream accepts a buffer and tokenizes the contents, then streams those
// tokens to the lexer.  The id is used to identify different files that are
// being tokenized at the same time (for includes).
func (tk Tokenizer) Stream(data []byte, len int, id int) {
	errors.NoImpl(tk.handler, "Tokenizer.Stream")
}

// BeginStream sets up buffers for the streaming process.  This method
// should be called from the source reading layer, and the filename should
// be as complete as possible since it is used for printing errors.
func (tk Tokenizer) BeginStream(filename string, id int) {
	errors.NoImpl(tk.handler, "Tokenizer.BeginStream")
}

// EndStream is called by the source reading layer once the entire file has
// been streamed into the tokenizer.
func (tk Tokenizer) EndStream(id int) {
	errors.NoImpl(tk.handler, "Tokenizer.EndStream")
}

// ReadFile starts the reading of a different source file.  This is used
// when reading an included file.  The id is used in the lexer to
// reconstruct all the file parts in order.  This call is proxied down to
// the source reader layer.
func (tk Tokenizer) ReadFile(filename string, id int, sourceType sourcereader.SourceType) {
	errors.NoImpl(tk.handler, "Tokenizer.ReadFile")
}
