package parser

// Equation represents a mathematical equation
type Equation struct {
	// LHS is the left-hand side of the equation
	LHS Expression
	// RHS is the right-hand side of the equation
	RHS Expression
}
