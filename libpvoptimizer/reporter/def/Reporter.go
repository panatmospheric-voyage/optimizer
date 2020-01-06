package def

import (
	libpvoptimizer "../.."
	"../../errors"
	"../../evaluator"
)

// Reporter is the default implementation of IReporter
type Reporter struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (rp Reporter) Init(evaluator libpvoptimizer.IEvaluator, resultwriter libpvoptimizer.IResultWriter, e errors.IErrorHandler) {
	rp.handler = e
	errors.NoImpl(rp.handler, "Reporter.Init")
}

// Report generates the report to save
func (rp Reporter) Report(model evaluator.OptimizedModel) {
	errors.NoImpl(rp.handler, "Reporter.Report")
}
