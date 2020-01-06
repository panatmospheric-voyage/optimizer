package parser

// Parameter represents a variable that the optimizer can change the value of
type Parameter struct {
	// Name is the name of the variable
	Name []string
	// Minimum is the lowest value the variable can have
	Minimum Number
	// MinimumInclude specifies if the minimum is inclusive
	MinimumInclude bool
	// Maximum is the highest value the variable can have
	Maximum Number
	// MaximumInclude specifies if the maximum is inclusive
	MaximumInclude bool
	// Unit is the unit the parameter has
	Unit Unit
	// Summarize defines if the parameter should be summarized
	Summarize bool
}
