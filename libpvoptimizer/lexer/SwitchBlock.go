package lexer

// SwitchBlock represents a single if or else/if block
type SwitchBlock struct {
	// LHS is the left-hand side of the condition
	LHS []ExpressionUnit
	// RHS is the right-hand side of the condition
	RHS []ExpressionUnit
	// Operator is the comparison operator between the LHS and RHS
	Operator ComparisonOperator
	// Statements are the statements to run if the branch runs
	Statements []Statement
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
