package lexer

// Unit represents one base unit in a composite unit
type Unit struct {
	// The name of the unit
	Name string
	// The exponent on the unit
	Power string
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
