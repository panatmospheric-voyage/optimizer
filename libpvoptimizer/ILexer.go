package libpvoptimizer

import "./tokenizer"
import "./errors"

// ILexer represents the interface for the lexing layer.  This is the third
// layer in the optimization pipeline and it handles the includes, makes sure
// tokens are in the correct order, and converts the tokens from strings to more
// usable data types (including merging tokens together).  After that, these
// lexemes are streamed into the parser layer.
type ILexer interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(tokenizer ITokenizer, parser IParser, e errors.IErrorHandler);
	// Stream accepts a token and processes it.  If it is the end of a
	// statement, it then streams it to the parser layer.
	Stream(token tokenizer.Token, id int);
	// EndStream is called by the tokenizer once one of the streams has finished
	// being tokenized and streamed into the lexer.
	EndStream(id int);
}
