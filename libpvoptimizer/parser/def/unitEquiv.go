package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func getBaseUnit(model *parser.Model, name string) *parser.BaseUnit {
	for i, u := range model.Units {
		if u.Name == name {
			return &model.Units[i]
		}
	}
	l := len(model.Units)
	model.Units = append(model.Units, parser.BaseUnit{
		Name: name,
	})
	return &model.Units[l]
}

func getUnit(model *parser.Model, name []lexer.Unit, e errors.IErrorHandler) parser.Unit {
	unit := parser.Unit{
		Parts: make([]parser.CompositeUnitPart, len(name)),
	}
	for i, u := range name {
		n, er := parser.ParseNumber(u.Power)
		if er != nil {
			e.Handle(errors.Error{
				Arguments: []interface{}{u.Power, er},
				Code:      errors.NumberParseError,
				LineNo:    u.LineNo,
				CharNo:    u.CharNo,
				FileName:  u.FileName,
			})
		}
		unit.Parts[i] = parser.CompositeUnitPart{
			Unit:  getBaseUnit(model, u.Name),
			Power: n,
		}
	}
	return unit
}

func parseUnitEquivalence(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	f, er := parser.ParseNumber(statement.Lexemes[3].Name)
	if er != nil {
		err(statement, e, errors.NumberParseError, statement.Lexemes[3].Name, er)
		return
	}
	model.UnitEquivalents = append(model.UnitEquivalents, parser.UnitEquivalence{
		Unit:           getBaseUnit(model, statement.Lexemes[1].Unit[0].Name),
		Factor:         f,
		EquivalentUnit: getUnit(model, statement.Lexemes[4].Unit, e),
	})
}

func parseUnitExactEquivalence(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	model.UnitEquivalents = append(model.UnitEquivalents, parser.UnitEquivalence{
		Unit:           getBaseUnit(model, statement.Lexemes[1].Unit[0].Name),
		Factor:         1,
		EquivalentUnit: getUnit(model, statement.Lexemes[3].Unit, e),
	})
}
