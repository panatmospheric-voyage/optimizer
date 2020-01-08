package def

import "../../parser"

func unitEquals(a, b parser.Unit) bool {
	if len(a.Parts) != len(b.Parts) {
		return false
	}
	for i, v := range a.Parts {
		if v.Unit != b.Parts[i].Unit || v.Power != b.Parts[i].Power {
			return false
		}
	}
	return true
}

func unitMultiply(a, b parser.Unit) parser.Unit {
	u := parser.Unit{
		Parts: make([]parser.CompositeUnitPart, len(a.Parts)),
	}
	copy(u.Parts, a.Parts)
	for _, p := range b.Parts {
		found := false
		for i, p2 := range u.Parts {
			if p2.Unit == p.Unit {
				u.Parts[i].Power += p.Power
				found = true
				break
			}
		}
		if !found {
			u.Parts = append(u.Parts, p)
		}
	}
	return u
}

func unitDivide(a, b parser.Unit) parser.Unit {
	u := parser.Unit{
		Parts: make([]parser.CompositeUnitPart, len(a.Parts)),
	}
	copy(u.Parts, a.Parts)
	for _, p := range b.Parts {
		found := false
		for i, p2 := range u.Parts {
			if p2.Unit == p.Unit {
				u.Parts[i].Power -= p.Power
				found = true
				break
			}
		}
		if !found {
			p.Power *= -1
			u.Parts = append(u.Parts, p)
		}
	}
	return u
}
