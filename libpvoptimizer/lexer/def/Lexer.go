package def

import (
	libpvoptimizer "../.."
	"../../errors"
	"../../tokenizer"
)

// Lexer is the default implementation of ILexer
type Lexer struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (lx Lexer) Init(tokenizer libpvoptimizer.ITokenizer, parser libpvoptimizer.IParser, e errors.IErrorHandler) {
	lx.handler = e
	errors.NoImpl(lx.handler, "Lexer.Init")
}

// Stream accepts a token and processes it.  If it is the end of a
// statement, it then streams it to the parser layer.
func (lx Lexer) Stream(token tokenizer.Token, id int) {
	errors.NoImpl(lx.handler, "Lexer.Stream")
}

// EndStream is called by the tokenizer once one of the streams has finished
// being tokenized and streamed into the lexer.
func (lx Lexer) EndStream(id int) {
	errors.NoImpl(lx.handler, "Lexer.EndStream")
}
