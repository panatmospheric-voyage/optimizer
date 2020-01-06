package lexer

// LexemeType represents a type of a lexeme
type LexemeType int;

const (
	// KeywordLiteral represents a language keyword or symbol
	KeywordLiteral LexemeType = 0
	// UnitLiteral represents the name of a unit
	UnitLiteral LexemeType = 1
	// NumberLiteral represents a literal number
	NumberLiteral LexemeType = 2
	// Expression represents a mathematical expression
	Expression LexemeType = 3
	// GroupBlock represents a block containing other statements
	GroupBlock LexemeType = 4
	// Identifier represents a literal string that represents a variable
	Identifier LexemeType = 5
	// Switch represents an if/else tree
	Switch LexemeType = 6
	// Conditional represents a conditional operator
	Conditional LexemeType = 7
)
