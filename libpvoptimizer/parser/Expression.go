package parser

import "../lexer"

// Expression represents a mathematical expression
type Expression struct {
	// Type is the type of expression
	Type ExpressionType
	// Value is the value of the constant
	Value Number
	// Unit is the unit of the constant
	Unit Unit
	// Name is the name of the variable
	Name []string
	// LHS is the left-hand side of the binary operator or the inside of the
	// function
	LHS *Expression
	// RHS is the right-hand side of the binary operator
	RHS *Expression
	// Function is the function to call
	Function lexer.Function
}
