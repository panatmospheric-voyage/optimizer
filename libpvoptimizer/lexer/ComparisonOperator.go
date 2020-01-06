package lexer

// ComparisonOperator represents a symbol that is used to compare two numbers
type ComparisonOperator int

const (
	// LessThan operator
	LessThan ComparisonOperator = 0
	// LessThanOrEqual operator
	LessThanOrEqual ComparisonOperator = 1
	// GreaterThan operator
	GreaterThan ComparisonOperator = 2
	// GreaterThanOrEqual operator
	GreaterThanOrEqual ComparisonOperator = 3
	// Equal operator
	Equal ComparisonOperator = 4
	// NotEqual operator
	NotEqual ComparisonOperator = 5
)
