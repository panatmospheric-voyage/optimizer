package evaluator

import "../parser"

// OptimizedModel represents the model after the optimizations are complete
type OptimizedModel struct {
	// FailedRequirements lists the requirements that could not be met
	FailedRequirements []parser.Requirement
	// Properties contains all the properties from the model
	Properties []Property
}
