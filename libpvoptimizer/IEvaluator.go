package libpvoptimizer

import "./parser"

// IEvaluator represents the interface for the evaluator layer.  This is the
// fifth layer in the optimization pipeline and it performs the actual
// optimizations on the model.  After that, the optimized model is sent to the
// reporter layer for a report to be generated.
type IEvaluator interface {
	// Init initializes the layer and is called from the pipeline layer
	Init(parser IParser, reporter IReporter);
	// Evaluate evaluates the model and optimizes it
	Evaluate(model parser.Model);
}
