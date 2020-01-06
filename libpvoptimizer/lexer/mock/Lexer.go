package mock

import (
	"fmt"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../sourcereader"
	"../../tokenizer"
)

// Lexer is the default implementation of ILexer
type Lexer struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (lx *Lexer) Init(tokenizer libpvoptimizer.ITokenizer, parser libpvoptimizer.IParser, e errors.IErrorHandler, wg *sync.WaitGroup) {
	lx.handler = e
	tokenizer.ReadFile("", 0, sourcereader.DefaultSource)
}

// Stream accepts a token and processes it.  If it is the end of a
// statement, it then streams it to the parser layer.
func (lx Lexer) Stream(token tokenizer.Token, id int) {
	fmt.Printf("Token: '%s' for %d at %s:%d:%d\n", token.Text, id, token.FileName, token.LineNo, token.CharNo)
}

// EndStream is called by the tokenizer once one of the streams has finished
// being tokenized and streamed into the lexer.
func (lx Lexer) EndStream(id int) {
	fmt.Printf("End of stream %d\n", id)
}
