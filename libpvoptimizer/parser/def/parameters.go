package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parseParameter(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	min, er := parser.ParseNumber(statement.Lexemes[3].Name)
	if er != nil {
		err(statement, e, errors.NumberParseError, statement.Lexemes[3].Name, er)
	}
	max, er := parser.ParseNumber(statement.Lexemes[5].Name)
	if er != nil {
		err(statement, e, errors.NumberParseError, statement.Lexemes[5].Name, er)
	}
	var u parser.Unit
	if len(statement.Lexemes) == 8 {
		u = getUnit(model, statement.Lexemes[7].Unit, e)
	} else {
		u = parser.Unit{
			Parts: []parser.CompositeUnitPart{},
		}
	}
	model.Parameters = append(model.Parameters, parser.Parameter{
		Name:           append(scope, statement.Lexemes[1].Name),
		Minimum:        min,
		MinimumInclude: statement.Lexemes[2].Keyword == lexer.InclusiveOpen,
		Maximum:        max,
		MaximumInclude: statement.Lexemes[6].Keyword == lexer.InclusiveClose,
		Unit:           u,
		Summarize:      false,
	})
}
