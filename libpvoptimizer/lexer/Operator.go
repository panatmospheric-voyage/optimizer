package lexer

// Operator represents a mathematical operation
type Operator int

const (
	// Addition operator
	Addition Operator = 0
	// Subtraction operator
	Subtraction Operator = 1
	// Multiplication operator
	Multiplication Operator = 2
	// Division operator
	Division Operator = 3
	// Exponentiation operator
	Exponentiation Operator = 4
)

var operatorString = []string{
	"Addition",
	"Subtraction",
	"Multiplication",
	"Division",
	"Exponentiation",
}

func (o Operator) String() string {
	if o >= 0 && int(o) < len(operatorString) {
		return operatorString[o]
	}
	return "Unknown Operator"
}
