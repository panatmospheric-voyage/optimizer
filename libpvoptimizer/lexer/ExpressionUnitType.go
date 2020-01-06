package lexer

// ExpressionUnitType represents the type of one token in an expression
type ExpressionUnitType int;

const (
	// ExpressionNumber is a literal number
	ExpressionNumber ExpressionUnitType = 0
	// Variable is an identifier or multiple identifiers separated by '.'
	Variable ExpressionUnitType = 1
	// OperatorSymbol is a mathematical operator symbol
	OperatorSymbol ExpressionUnitType = 2
	// FunctionLiteral is a function
	FunctionLiteral ExpressionUnitType = 3
)
