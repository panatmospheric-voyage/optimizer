package mock

import (
	"fmt"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
)

// ResultWriter is the default implementation of IResultWriter
type ResultWriter struct {
}

// Init initializes the layer and is called from the pipeline layer
func (rw ResultWriter) Init(reporter libpvoptimizer.IReporter, e errors.IErrorHandler, wg *sync.WaitGroup) {
}

// Save the results
func (rw ResultWriter) Save(results string) {
	fmt.Print(results)
}

// SetOutputFile sets the file the results will be saved to.  If not specified,
// the results will be printed to standard output.
func (rw ResultWriter) SetOutputFile(filename string) {
	fmt.Printf("Saving output to %s\n", filename)
}
