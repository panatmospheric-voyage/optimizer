package parser

// Property is a variable that is set by an equation or constant
type Property struct {
	// Name is the name of the property
	Name []string
	// Definition is the equation defining the property
	Definition Equation
	// Summarize defines if the property should be summarized
	Summarize bool
}
