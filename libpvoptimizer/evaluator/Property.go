package evaluator

import "../parser"

// Property represents a property that has been optimized
type Property struct {
	// Name is the name of the property or enum
	Name []string
	// EnumValue is the value of the enum
	EnumValue string
	// Value is the value of the property
	Value parser.Number
	// Unit is the units on the property value
	Unit parser.Unit
	// Summarize defines if the property should be summarized
	Summarize bool
}
