package parser

// Model represents the entire model to be optimized
type Model struct {
	// Units contains all the units defined in the model
	Units []BaseUnit
	// UnitEquivalents contains all the unit equivalences defined in the file
	UnitEquivalents []UnitEquivalence
	// Parameters contains all the variables the optimizer can change
	Parameters []Parameter
	// UniversalProperties contains all properties that are the same for all
	// enumerations
	UniversalProperties []Property
	// Enumerations contains all properties that can change discretely
	Enumerations []Enumeration
	// Requirements contains all property requirements
	Requirements []Requirement
	// Optimization options
	Optimization Optimization
}
