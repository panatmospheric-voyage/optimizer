package parser

import (
	"fmt"
	"strings"
)

// Unit represents a composite unit
type Unit struct {
	// Parts contains all of the parts of the composite unit
	Parts []CompositeUnitPart
}

func (u Unit) String() string {
	strs := make([]string, len(u.Parts))
	for i, p := range u.Parts {
		if p.Power == 1 {
			strs[i] = p.Unit.Name
		} else {
			strs[i] = fmt.Sprintf("%s^%s", p.Unit.Name, p.Power)
		}
	}
	return strings.Join(strs, "*")
}
