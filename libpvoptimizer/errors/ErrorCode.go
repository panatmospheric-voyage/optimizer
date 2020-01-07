package errors

// ErrorCode is the code for an error
type ErrorCode int

var strings []string = []string{
	"Function '%s' not implemented.",
	"Error reading file: %s.",
	"Invalid state %d in state machine %s.",
	"Illegal empty statement.",
	"Missing case in compiler.",
	"Expected statement but got '%s'.",
	"Expected unit but got '%s'.",
	"Expected '=' but got '%s'.",
	"Unexpected 'else' not after 'if'.",
	"Expected number or unit but got '%s'.",
	"Expected ';' but got '%s'.",
	"Expected identifier but got '%s'.",
	"Expected '{' but got '%s'.",
	"Expected '(' or '[' but got '%s'.",
	"Expected ')' or ']' but got '%s'.",
	"Expected ',' but got '%s'.",
	"Expected number but got '%s'.",
	"Unexpected identifier (got '%s').",
	"Expected '<=', '<', '>=', '>', '==', or '!=' but got '%s'.",
	"Expected '^' or '*' but got '%s'.",
	"Expected '*' but got '%s'.",
	"Expected operator but got '%s'.",
	"Expected expression but got '%s'.",
	"Expected '(' but got '%s'.",
	"Unexpected '}' in top-level scope.",
}

const (
	// NotImplemented error
	NotImplemented ErrorCode = 0
	// FileReadError error
	FileReadError ErrorCode = 1
	// StateError error
	StateError ErrorCode = 2
	// EmptyStatement error
	EmptyStatement ErrorCode = 3
	// MissingCase error
	MissingCase ErrorCode = 4
	// ExpectedStatement error
	ExpectedStatement ErrorCode = 5
	// ExpectedUnit error
	ExpectedUnit ErrorCode = 6
	// ExpectedEquals error
	ExpectedEquals ErrorCode = 7
	// UnexpectedElse error
	UnexpectedElse ErrorCode = 8
	// ExpectedNumberOrUnit error
	ExpectedNumberOrUnit ErrorCode = 9
	// ExpectedEndOfStatement error
	ExpectedEndOfStatement ErrorCode = 10
	// ExpectedIdentifier error
	ExpectedIdentifier ErrorCode = 11
	// ExpectedBlockStart error
	ExpectedBlockStart ErrorCode = 12
	// ExpectedRangeOpen error
	ExpectedRangeOpen ErrorCode = 13
	// ExpectedRangeClose error
	ExpectedRangeClose ErrorCode = 14
	// ExpectedDelimiter error
	ExpectedDelimiter ErrorCode = 15
	// ExpectedNumber error
	ExpectedNumber ErrorCode = 16
	// UnexpectedIdentifier error
	UnexpectedIdentifier ErrorCode = 17
	// ExpectedCondition error
	ExpectedCondition ErrorCode = 18
	// ExpectedExpMul error
	ExpectedExpMul ErrorCode = 19
	// ExpectedMul error
	ExpectedMul ErrorCode = 20
	// ExpectedOperator error
	ExpectedOperator ErrorCode = 21
	// ExpectedExpression error
	ExpectedExpression ErrorCode = 22
	// ExpectedParenthesis error
	ExpectedParenthesis ErrorCode = 23
	// UnexpectedEndBlock error
	UnexpectedEndBlock ErrorCode = 24
)

// String returns the format string for an error
func (e ErrorCode) String() string {
	return strings[e]
}
