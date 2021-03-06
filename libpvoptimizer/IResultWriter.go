package libpvoptimizer

import (
	"sync"

	"./errors"
)

// IResultWriter represents the interface for the result writing layer.  This is
// the seventh layer in the optimization pipeline and it saves the results.
type IResultWriter interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(reporter IReporter, e errors.IErrorHandler, wg *sync.WaitGroup)
	// Save the results
	Save(results string)
}
