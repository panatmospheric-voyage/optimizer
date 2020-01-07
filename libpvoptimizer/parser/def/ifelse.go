package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parseIfElse(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parseIfElse")
}
