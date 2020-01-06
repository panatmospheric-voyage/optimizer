package lexer

// Lexeme represents a single token or group of tokens that have one meaning.
type Lexeme struct {
	// Type is the type of this lexeme
	Type LexemeType
	// Name is the text content of a literal
	Name string
	// Unit is the units in the lexeme
	Unit []Unit
	// Expression is the expression the lexeme represents
	Expression []ExpressionUnit
	// Statements are the statements inside the group block
	Statements []Statement
	// SwitchBlocks are the parts of the if/else tree
	SwitchBlocks []SwitchBlock
	// The keyword the lexeme represents
	Keyword Keyword
	// Condition represents the comparison operator this lexeme is
	Condition ComparisonOperator
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
