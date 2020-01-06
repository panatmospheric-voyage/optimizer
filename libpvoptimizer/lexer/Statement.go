package lexer

// Statement represents a group of lexemes that together form a single statement
// in the source code
type Statement struct {
	// Lexemes are the lexemes that make up the statement
	Lexemes []Lexeme
}
