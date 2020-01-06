package def

import (
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../lexer"
)

// Parser is the default implementation of IParser
type Parser struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (ps Parser) Init(lexer libpvoptimizer.ILexer, evaluator libpvoptimizer.IEvaluator, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ps.handler = e
	errors.NoImpl(ps.handler, "Parser.Init")
}

// Stream accepts a statement and processes it.
func (ps Parser) Stream(statement lexer.Statement) {
	errors.NoImpl(ps.handler, "Parser.Stream")
}

// End is called by the lexer once all statements have been streamed into
// the parser.  This is when the model is sent into the evaluator layer.
func (ps Parser) End() {
	errors.NoImpl(ps.handler, "Parser.End")
}
