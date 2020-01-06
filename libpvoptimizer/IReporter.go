package libpvoptimizer

import (
	"sync"

	"./errors"
	"./evaluator"
)

// IReporter represents the interface for the reporter layer.  This is the sixth
// layer in the optimization pipeline and it generates the report from the
// optimized model.  After that, the report is sent to the result writing layer
// to be saved.
type IReporter interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(evaluator IEvaluator, resultwriter IResultWriter, e errors.IErrorHandler, wg *sync.WaitGroup)
	// Report generates the report to save
	Report(model evaluator.OptimizedModel)
}
