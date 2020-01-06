package tokenizer

// Token represents one token from the source file
type Token struct {
	// Text contains the actual content of the token
	Text string
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}
