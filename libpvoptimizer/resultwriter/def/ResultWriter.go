package def

import (
	"fmt"
	"io/ioutil"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
)

// ResultWriter is the default implementation of IResultWriter
type ResultWriter struct {
	handler    errors.IErrorHandler
	outputFile string
}

// Init initializes the layer and is called from the pipeline layer
func (rw *ResultWriter) Init(reporter libpvoptimizer.IReporter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	rw.handler = e
}

// Save the results
func (rw ResultWriter) Save(results string) {
	if rw.outputFile == "" {
		fmt.Print(results)
	} else {
		if err := ioutil.WriteFile(rw.outputFile, []byte(results), 0); err != nil {
			rw.handler.Handle(errors.Error{
				Arguments: []interface{}{err.Error()},
				Code:      errors.FileWriteError,
				LineNo:    -1,
				CharNo:    -1,
				FileName:  rw.outputFile,
			})
		}
	}
}

// SetOutputFile sets the file the results will be saved to.  If not specified,
// the results will be printed to standard output.
func (rw *ResultWriter) SetOutputFile(filename string) {
	rw.outputFile = filename
}
