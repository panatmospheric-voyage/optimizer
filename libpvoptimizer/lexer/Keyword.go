package lexer

// Keyword represents which keyword a lexeme is
type Keyword int

const (
	// UnitKeyword keyword
	UnitKeyword = 0
	// EqualsKeyword keyword
	EqualsKeyword = 1
	// PropertyKeyword keyword
	PropertyKeyword = 2
	// AssemblyKeyword keyword
	AssemblyKeyword = 3
	// ParameterKeyword keyword
	ParameterKeyword = 4
	// SummarizeKeyword keyword
	SummarizeKeyword = 5
	// EnumKeyword keyword
	EnumKeyword = 6
	// ValueKeyword keyword
	ValueKeyword = 7
	// InclusiveOpen keyword
	InclusiveOpen = 8
	// ExclusiveOpen keyword
	ExclusiveOpen = 9
	// InclusiveClose keyword
	InclusiveClose = 10
	// ExclusiveClose keyword
	ExclusiveClose = 11
	// AssignArrow keyword
	AssignArrow = 12
	// RequireKeyword keyword
	RequireKeyword = 13
	// MinimizeKeyword keyword
	MinimizeKeyword = 14
	// MaximizeKeyword keyword
	MaximizeKeyword = 15
	// CommaKeyword keyword
	CommaKeyword = 16
	// ZeroKeyword keyword
	ZeroKeyword = 17
)

var keywordString = []string{
	"UnitKeyword",
	"EqualsKeyword",
	"PropertyKeyword",
	"AssemblyKeyword",
	"ParameterKeyword",
	"SummarizeKeyword",
	"EnumKeyword",
	"ValueKeyword",
	"InclusiveOpen",
	"ExclusiveOpen",
	"InclusiveClose",
	"ExclusiveClose",
	"AssignArrow",
	"RequireKeyword",
	"MinimizeKeyword",
	"MaximizeKeyword",
	"CommaKeyword",
	"ZeroKeyword",
}

func (k Keyword) String() string {
	if k >= 0 && int(k) < len(keywordString) {
		return keywordString[k]
	}
	return "Unknown Keyword"
}
