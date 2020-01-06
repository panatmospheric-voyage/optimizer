package def

import (
	"io"
	"os"
	"path"
	"sync"

	sourcereader ".."
	libpvoptimizer "../.."
	"../../errors"

	"github.com/golang-collections/collections/stack"
)

// SourceReader is the default implementation of ISourceReader
type SourceReader struct {
	handler     errors.IErrorHandler
	defaultFile string
	tokenizer   libpvoptimizer.ITokenizer
	fileStack   *stack.Stack
	wg          *sync.WaitGroup
}

type fileStackEntry struct {
	file *os.File
	id   int
	name string
}

func (sr *SourceReader) readError(err error, filename string) {
	sr.handler.Handle(errors.Error{
		Arguments: []interface{}{err.Error()},
		Code:      errors.FileReadError,
		LineNo:    -1,
		CharNo:    -1,
		FileName:  filename,
	})
}

func (sr *SourceReader) readThread() {
	defer sr.wg.Done()
	buf := make([]byte, 4096)
	for sr.fileStack.Len() > 0 {
		ent := sr.fileStack.Peek().(fileStackEntry)
		n, err := ent.file.Read(buf)
		if err != nil && err != io.EOF {
			sr.readError(err, ent.name)
			sr.fileStack.Pop()
			if err = ent.file.Close(); err != nil {
				sr.readError(err, ent.name)
			}
		}
		if n == 0 {
			sr.tokenizer.EndStream(ent.id)
			sr.fileStack.Pop()
			if ent.file != os.Stdin {
				if err = ent.file.Close(); err != nil {
					sr.readError(err, ent.name)
				}
			}
		} else {
			sr.tokenizer.Stream(buf, n, ent.id)
		}
	}
}

// Init initializes the layer and is called from the pipeline layer
func (sr *SourceReader) Init(tokenizer libpvoptimizer.ITokenizer, e errors.IErrorHandler, wg *sync.WaitGroup) {
	sr.handler = e
	sr.tokenizer = tokenizer
	sr.fileStack = stack.New()
	sr.wg = wg
}

// ReadFile starts the reading of a different source file.  This is used
// when reading an included file.  The id is used in the lexer to
// reconstruct all the file parts in order.
func (sr *SourceReader) ReadFile(filename string, id int, sourceType sourcereader.SourceType) {
	var file *os.File = nil
	var name string
	switch sourceType {
	case sourcereader.DefaultSource:
		if sr.defaultFile == "" {
			file = os.Stdin
			name = "stdin"
		} else {
			name = sr.defaultFile
		}
		break
	case sourcereader.UserSource:
		name = filename
		break
	case sourcereader.SystemSource:
		name = path.Join("/usr/include/pvoptimizer/", filename)
		break
	default:
		return
	}
	if file == nil {
		var err error
		if file, err = os.Open(name); err != nil {
			sr.readError(err, name)
			return
		}
	}
	sr.tokenizer.BeginStream(name, id)
	sr.fileStack.Push(fileStackEntry{
		file: file,
		id:   id,
		name: name,
	})
	if sr.fileStack.Len() == 1 {
		sr.wg.Add(1)
		go sr.readThread()
	}
}

// SetDefaultFile sets the file that will be opened first.  If not specified,
// the first file is read from standard input.
func (sr *SourceReader) SetDefaultFile(filename string) {
	sr.defaultFile = filename
}
