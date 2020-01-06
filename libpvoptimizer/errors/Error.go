package errors

import "fmt"

// Error represents an error in the code
type Error struct {
	// Arguments to the string format for the error
	Arguments []string
	// Code of error
	Code ErrorCode
	// LineNo contains the line number the token was found on
	LineNo int
	// CharNo contains the character number in the line of the start of the
	// token
	CharNo int
	// FileName is the name of the file the token was found in
	FileName string
}

// FormatError returns the formatted error description
func (e Error) FormatError() string {
	return fmt.Sprintf(e.Code.String(), e.Arguments)
}
