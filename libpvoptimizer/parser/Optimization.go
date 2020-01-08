package parser

// Optimization represents the options for the optimization
type Optimization struct {
	// Type of optimization to perform
	Type OptimizationType
	// Variable to optimize
	Variable []string
	// Minimum is the smallest change that needs to be observed in any variables
	Minimum Number
	// Accuracy to solve to
	Accuracy Number
	// Iterations is the number of times to run the gradient descent algorithm
	Iterations int
	// Seed for the random number generator
	Seed int64
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
