package def

import (
	"fmt"

	lexer ".."
	"../../errors"
	"../../tokenizer"
)

type unitReader struct {
	handler  errors.IErrorHandler
	state    unitState
	target   *[]lexer.Unit
	inverted bool
}

func (ur *unitReader) Init(handler errors.IErrorHandler) {
	ur.handler = handler
}

func (ur *unitReader) Reset(target *[]lexer.Unit) {
	ur.state = readUnitName
	ur.target = target
	ur.inverted = false
}

func (ur *unitReader) err(token tokenizer.Token, code errors.ErrorCode, args ...interface{}) {
	ur.handler.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    token.LineNo,
		CharNo:    token.CharNo,
		FileName:  token.FileName,
	})
}

func (ur *unitReader) Read(token tokenizer.Token) sublexResult {
	switch ur.state {
	case readUnitName:
		if isIdentifier(token.Text) {
			*ur.target = append(*ur.target, lexer.Unit{
				Name:     token.Text,
				Power:    "1",
				LineNo:   token.LineNo,
				CharNo:   token.CharNo,
				FileName: token.FileName,
			})
			ur.state = afterUnitName
		} else {
			ur.err(token, errors.ExpectedIdentifier, token.Text)
			return slError
		}
		break
	case afterUnitName:
		switch token.Text {
		case "^":
			ur.state = readUnitPower
			break
		case "*":
			ur.state = readUnitName
			break
		case "/":
			if ur.inverted {
				ur.err(token, errors.ExpectedExpMul, "/")
				return slError
			}
			ur.inverted = true
			ur.state = readUnitName
			break
		default:
			return slComplete
		}
		break
	case readUnitPower:
		u := &(*ur.target)[len(*ur.target)-1]
		if ur.inverted {
			if token.Text[0] == '-' {
				u.Power = token.Text[1:]
			} else {
				u.Power = fmt.Sprintf("-%s", token.Text)
			}
		} else {
			u.Power = token.Text
		}
		ur.state = afterUnitPower
		break
	case afterUnitPower:
		switch token.Text {
		case "*":
			ur.state = readUnitName
			break
		case "/":
			if ur.inverted {
				ur.err(token, errors.ExpectedMul, "/")
				return slError
			}
			ur.inverted = true
			ur.state = readUnitName
			break
		default:
			return slComplete
		}
		break
	default:
		ur.err(token, errors.StateError, ur.state, "unitState")
		ur.state = readUnitName
		break
	}
	return slValid
}
