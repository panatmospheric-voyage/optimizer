package lexer

// ExpressionUnit represents one token in an expression
type ExpressionUnit struct {
	// Type is the type of the expression unit
	Type ExpressionUnitType
	// Text is the value of the number or variable (each element is an
	// identifier separated by '.' if it is a variable, otherwise there is only
	// one element)
	Text []string
	// Unit is the unit the number has
	Unit []Unit
	// Operator that the expression is
	Operator Operator
	// Function that the expression is
	Function Function
	// SubExpression is the expression inside a function parentheses
	SubExpression []ExpressionUnit
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
