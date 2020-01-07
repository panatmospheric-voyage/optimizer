package def

import (
	"strconv"
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
	l := statement.Lexemes[len(statement.Lexemes)-1]
	if len(l.Expression) != 1 || l.Expression[0].Type != lexer.Variable {
		err(statement, e, errors.UnaryExpectedVariable, stName)
		return nil, nil, []string{}
	}
	name := append(scope, l.Expression[0].Text...)
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
	if model.Optimization.Type != parser.NoOptimize {
		err(statement, e, errors.MultipleOptimizations)
	} else {
		prop, param, name := unaryGetObject(model, statement, scope, e, stName)
		if prop == nil {
			if param != nil {
				err(statement, e, errors.OptimizeParameter, strings.Join(name, "."))
			}
		} else {
			mean, er := parser.ParseNumber(statement.Lexemes[2].Name)
			if er != nil {
				err(statement, e, errors.NumberParseError, statement.Lexemes[2].Name, er)
				mean = 1
			}
			acc, er := parser.ParseNumber(statement.Lexemes[4].Name)
			if er != nil {
				err(statement, e, errors.NumberParseError, statement.Lexemes[4].Name, er)
				acc = 0.001
			}
			var iters int64 = 0
			var seed int64 = 0
			if len(statement.Lexemes) >= 9 {
				iters, er = strconv.ParseInt(statement.Lexemes[6].Name, 10, 32)
				if er != nil {
					err(statement, e, errors.NumberParseError, statement.Lexemes[6].Name, er)
					iters = 0
				}
				if len(statement.Lexemes) == 11 {
					seed, er = strconv.ParseInt(statement.Lexemes[8].Name, 10, 64)
					if er != nil {
						err(statement, e, errors.NumberParseError, statement.Lexemes[8].Name, er)
						seed = 0
					}
				}
			}
			model.Optimization = parser.Optimization{
				Type:       t,
				Variable:   name,
				Mean:       mean,
				Accuracy:   acc,
				Iterations: int(iters),
				Seed:       seed,
			}
		}
	}
}

func parseMinimize(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	parseOptimization(model, statement, scope, e, "minimize", parser.Minimization)
}

func parseMaximize(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	parseOptimization(model, statement, scope, e, "maximize", parser.Maximization)
}

func parseZero(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	parseOptimization(model, statement, scope, e, "zero", parser.Zero)
}
