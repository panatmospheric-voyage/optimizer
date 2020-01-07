package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parseRequirement(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	errors.NoImpl(e, "parseRequirement")
}
