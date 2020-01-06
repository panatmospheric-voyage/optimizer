package parser

import "../lexer"

// Requirement represents a requirement for a variable
type Requirement struct {
	// Name is the variable for the requirement
	Name []string
	// Condition is the requirement type
	Condition lexer.ComparisonOperator
	// Value is the expression defining the condition
	Value Expression
}
