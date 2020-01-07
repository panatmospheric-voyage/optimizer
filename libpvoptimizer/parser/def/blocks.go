package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parseBlockAssembly(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	sc := append(scope, statement.Lexemes[0].Name)
	for _, st := range statement.Lexemes[0].Statements {
		stream(model, st, sc, e)
	}
}

func parseBlockEnum(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parseBlockEnum")
}

func parseBlockValue(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parseBlockValue")
}

func parseBlock(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	switch statement.Lexemes[0].Keyword {
	case lexer.AssemblyKeyword:
		parseBlockAssembly(model, statement, scope, e)
		break
	case lexer.EnumKeyword:
		parseBlockEnum(model, statement, scope, e)
		break
	case lexer.ValueKeyword:
		parseBlockValue(model, statement, scope, e)
		break
	default:
		err(statement, e, errors.MissingCase)
		break
	}
}
