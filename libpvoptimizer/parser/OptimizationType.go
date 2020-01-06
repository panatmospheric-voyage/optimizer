package parser

// OptimizationType represents the type of optimization that is occuring
type OptimizationType int;

const (
	// Minimization optimization
	Minimization OptimizationType = 0
	// Maximization optimization
	Maximization OptimizationType = 1
)
