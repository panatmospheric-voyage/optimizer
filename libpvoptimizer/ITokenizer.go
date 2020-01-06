package libpvoptimizer

import (
	"./errors"
	"./sourcereader"
)

// ITokenizer represents the interface for the tokenizing layer.  This is the
// second layer in the optimization pipeline and it splits the data streamed
// from the source reader into tokens.
type ITokenizer interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(sourcereader ISourceReader, lexer ILexer, e errors.IErrorHandler)
	// Stream accepts a buffer and tokenizes the contents, then streams those
	// tokens to the lexer.  The id is used to identify different files that are
	// being tokenized at the same time (for includes).
	Stream(data []byte, len int, id int)
	// BeginStream sets up buffers for the streaming process.  This method
	// should be called from the source reading layer, and the filename should
	// be as complete as possible since it is used for printing errors.
	BeginStream(filename string, id int)
	// EndStream is called by the source reading layer once the entire file has
	// been streamed into the tokenizer.
	EndStream(id int)
	// ReadFile starts the reading of a different source file.  This is used
	// when reading an included file.  The id is used in the lexer to
	// reconstruct all the file parts in order.  This call is proxied down to
	// the source reader layer.
	ReadFile(filename string, id int, sourceType sourcereader.SourceType)
}
