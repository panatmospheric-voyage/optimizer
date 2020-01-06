package errors

// ErrorCode is the code for an error
type ErrorCode int;

var strings []string = []string {
	"Function '%s' not implemented.",
}

const (
	// NotImplemented error
	NotImplemented ErrorCode = 0
)

// String returns the format string for an error
func (e ErrorCode) String() string {
	return strings[e];
}
