package parser

// UnitEquivalence represents a binding between two units
type UnitEquivalence struct {
	// Unit is the unit being defined
	Unit BaseUnit
	// Factor is the number to multiply the EquivalentUnit by to get a Unit
	Factor Number
	// EquivalentUnit is the already-defined unit
	EquivalentUnit Unit
}
