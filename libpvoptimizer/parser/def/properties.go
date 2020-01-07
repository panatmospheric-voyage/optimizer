package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parsePropertyPrototype(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parsePropertyPrototype")
}

func parsePropertyDefinition(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	name := append(scope, statement.Lexemes[1].Name)
	model.UniversalProperties = append(model.UniversalProperties, parser.Property{
		Name: name,
		Definition: parser.Equation{
			LHS: parser.Expression{
				Type: parser.Variable,
				Name: name,
			},
			RHS: parseExpression(statement.Lexemes[3].Expression, model, e),
		},
		Summarize: false,
	})
}

func parsePropertyAssignment(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parsePropertyAssignment")
}
