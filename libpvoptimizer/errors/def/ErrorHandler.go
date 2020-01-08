package def

import (
	"fmt"
	"os"

	errors ".."
)

// ErrorHandler is the default implementation of the IErrorHandler
type ErrorHandler struct{}

// Handle the error
func (ErrorHandler) Handle(e errors.Error) {
	if e.Code == errors.MissingCase {
		fmt.Print("")
	}
	if e.FileName != "" {
		if e.LineNo >= 0 {
			if e.CharNo >= 0 {
				fmt.Fprintf(os.Stderr, "Error %d at %s:%d:%d\n", e.Code, e.FileName, e.LineNo, e.CharNo)
			} else {
				fmt.Fprintf(os.Stderr, "Error %d at %s:%d\n", e.Code, e.FileName, e.LineNo)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error %d in %s\n", e.Code, e.FileName)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error %d\n", e.Code)
	}
	fmt.Fprintln(os.Stderr, e.FormatError())
	fmt.Fprintln(os.Stderr, "")
}
