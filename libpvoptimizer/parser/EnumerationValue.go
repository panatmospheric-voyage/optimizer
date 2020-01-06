package parser

// EnumerationValue represents one discrete set of values an enumeration can
// have
type EnumerationValue struct {
	// Name is the name of this value
	Name string
	// Properties are the properties for this value
	Properties []Property
}
