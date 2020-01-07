package def

import (
	"strings"

	parser ".."
	"../../errors"
	"../../lexer"
)

func getNamedObject(model *parser.Model, name []string) (*parser.Property, *parser.Parameter) {
	for i, p := range model.UniversalProperties {
		if equals(p.Name, name) {
			return &model.UniversalProperties[i], nil
		}
	}
	for i, p := range model.Parameters {
		if equals(p.Name, name) {
			return nil, &model.Parameters[i]
		}
	}
	return nil, nil
}

func unaryGetObject(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler, stName string) (*parser.Property, *parser.Parameter, []string) {
	if len(statement.Lexemes[1].Expression) != 1 || statement.Lexemes[1].Expression[0].Type != lexer.Variable {
		err(statement, e, errors.UnaryExpectedVariable, stName)
		return nil, nil, []string{}
	}
	name := append(scope, statement.Lexemes[1].Expression[0].Text...)
	prop, param := getNamedObject(model, name)
	if prop == nil && param == nil {
		err(statement, e, errors.UnknownVariable, strings.Join(name, "."))
	}
	return prop, param, name
}

func parseSummarize(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	prop, param, _ := unaryGetObject(model, statement, scope, e, "summarize")
	if prop != nil {
		prop.Summarize = true
	}
	if param != nil {
		param.Summarize = true
	}
}

func parseOptimization(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler, stName string, t parser.OptimizationType) {
	if model.Type != parser.NoOptimize {
		err(statement, e, errors.MultipleOptimizations)
	} else {
		prop, param, name := unaryGetObject(model, statement, scope, e, stName)
		if prop == nil {
			if param != nil {
				err(statement, e, errors.OptimizeParameter, strings.Join(name, "."))
			}
		} else {
			model.Type = t
			model.Variable = name
		}
	}
}

func parseMinimize(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	parseOptimization(model, statement, scope, e, "minimize", parser.Minimization)
}

func parseMaximize(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	parseOptimization(model, statement, scope, e, "maximize", parser.Maximization)
}
