package libpvoptimizer

import (
	"sync"

	"./errors"
	"./lexer"
)

// IParser represents the interface for the parsing layer.  This is the fourth
// layer in the optimization pipeline and it converts all of the statements from
// the lexer into usable data structures.  After that, it sends the finished
// model to the evaluator to optimize it.
type IParser interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(lexer ILexer, evaluator IEvaluator, e errors.IErrorHandler, wg *sync.WaitGroup)
	// Stream accepts a statement and processes it.
	Stream(statement lexer.Statement)
	// End is called by the lexer once all statements have been streamed into
	// the parser.  This is when the model is sent into the evaluator layer.
	End()
}
