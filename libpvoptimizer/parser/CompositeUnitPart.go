package parser

// CompositeUnitPart represents part of a composite unit
type CompositeUnitPart struct {
	// Unit is the unit
	Unit *BaseUnit
	// Power is the exponent of the unit
	Power Number
}
