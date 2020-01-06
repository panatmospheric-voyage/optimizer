package parser

// Enumeration represents a set of properties that can change values discretely
type Enumeration struct {
	// Name is the name of the enumeration
	Name []string
	// Values contains the possible values of the enumeration
	Values []EnumerationValue
	// Summarize defines if the enumeration should be summarized
	Summarize bool
}
