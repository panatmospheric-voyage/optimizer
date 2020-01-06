package parser

// ExpressionType is the type of an expression
type ExpressionType int

const (
	// Constant value
	Constant ExpressionType = 0
	// Variable name
	Variable ExpressionType = 1
	// Addition operator
	Addition ExpressionType = 2
	// Subtraction operator
	Subtraction ExpressionType = 3
	// Multiplication operator
	Multiplication ExpressionType = 4
	// Division operator
	Division ExpressionType = 5
	// Exponentiation operator
	Exponentiation ExpressionType = 6
	// Function call
	Function ExpressionType = 7
)
