package def

import (
	sourcereader ".."
	libpvoptimizer "../.."
	"../../errors"
)

// SourceReader is the default implementation of ISourceReader
type SourceReader struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (sr SourceReader) Init(tokenizer libpvoptimizer.ITokenizer, e errors.IErrorHandler) {
	sr.handler = e
	errors.NoImpl(sr.handler, "SourceReader.Init")
}

// ReadFile starts the reading of a different source file.  This is used
// when reading an included file.  The id is used in the lexer to
// reconstruct all the file parts in order.
func (sr SourceReader) ReadFile(filename string, id int, sourceType sourcereader.SourceType) {
	errors.NoImpl(sr.handler, "SourceReader.ReadFile")
}

// SetDefaultFile sets the file that will be opened first.  If not specified,
// the first file is read from standard input.
func (sr SourceReader) SetDefaultFile(filename string) {
	errors.NoImpl(sr.handler, "SourceReader.SetDefaultFile")
}
