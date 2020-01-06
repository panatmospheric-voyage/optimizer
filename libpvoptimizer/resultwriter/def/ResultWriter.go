package def

import (
	libpvoptimizer "../.."
	"../../errors"
)

// ResultWriter is the default implementation of IResultWriter
type ResultWriter struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (rw ResultWriter) Init(reporter libpvoptimizer.IReporter, e errors.IErrorHandler) {
	rw.handler = e
	errors.NoImpl(rw.handler, "ResultWriter.Init")
}

// Save the results
func (rw ResultWriter) Save(results string) {
	errors.NoImpl(rw.handler, "ResultWriter.Save")
}

// SetOutputFile sets the file the results will be saved to.  If not specified,
// the results will be printed to standard output.
func (rw ResultWriter) SetOutputFile(filename string) {
	errors.NoImpl(rw.handler, "ResultWriter.SetOutputFile")
}
