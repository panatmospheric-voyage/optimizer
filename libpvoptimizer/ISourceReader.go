package libpvoptimizer

import "./sourcereader"
import "./errors"

// ISourceReader represents the interface for the source reading layer.  This is
// the first layer in the optimization pipeline and it reads the source files
// from the filesystem (in the default implementation).
type ISourceReader interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(tokenizer ITokenizer, e errors.IErrorHandler);
	// ReadFile starts the reading of a different source file.  This is used
	// when reading an included file.  The id is used in the lexer to
	// reconstruct all the file parts in order.
	ReadFile(filename string, id int, sourceType sourcereader.SourceType);
}
