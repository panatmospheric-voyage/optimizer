package parser

// Optimization represents the options for the optimization
type Optimization struct {
	// Type of optimization to perform
	Type OptimizationType
	// Variable to optimize
	Variable []string
	// Mean is the positive mean of optimization variable
	Mean Number
	// Accuracy to solve to
	Accuracy Number
	// Iterations is the number of times to run the gradient descent algorithm
	Iterations int
	// Seed for the random number generator
	Seed int64
}
