package def

import (
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../parser"
)

// Evaluator is the default implementation of IEvaluator
type Evaluator struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (ev Evaluator) Init(parser libpvoptimizer.IParser, reporter libpvoptimizer.IReporter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ev.handler = e
	errors.NoImpl(ev.handler, "Evaluator.Init")
}

// Evaluate evaluates the model and optimizes it
func (ev Evaluator) Evaluate(model parser.Model) {
	errors.NoImpl(ev.handler, "Evaluator.Evaluate")
}
