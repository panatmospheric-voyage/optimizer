package parser

// OptimizationType represents the type of optimization that is occuring
type OptimizationType int

const (
	// NoOptimize optimization
	NoOptimize OptimizationType = 0
	// Minimization optimization
	Minimization OptimizationType = 1
	// Maximization optimization
	Maximization OptimizationType = 2
)
