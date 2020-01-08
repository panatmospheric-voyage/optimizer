package def

import (
	"../../errors"
	"../../parser"
)

func solve(eq *parser.Equation, name []string, e errors.IErrorHandler) bool {
	if eq.LHS.Type == parser.Variable && equals(name, eq.LHS.Name) {
		return true
	}
	errors.NoImpl(e, "solve")
	return false
}
