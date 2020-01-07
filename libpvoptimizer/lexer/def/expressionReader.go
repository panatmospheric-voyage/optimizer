package def

import (
	lexer ".."
	"../../errors"
	"../../tokenizer"
	"github.com/golang-collections/collections/stack"
)

type expressionReader struct {
	handler   errors.IErrorHandler
	exprStack *stack.Stack
	state     expressionState
	unit      unitReader
}

func (ex *expressionReader) Init(handler errors.IErrorHandler) {
	ex.handler = handler
	ex.unit.Init(handler)
}

func (ex *expressionReader) Reset(target *[]lexer.ExpressionUnit) {
	ex.exprStack = stack.New()
	ex.exprStack.Push(target)
	ex.state = readVariable
}

func (ex *expressionReader) err(token tokenizer.Token, code errors.ErrorCode, args ...interface{}) {
	ex.handler.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    token.LineNo,
		CharNo:    token.CharNo,
		FileName:  token.FileName,
	})
}

func (ex *expressionReader) readOperator(token tokenizer.Token, e **lexer.ExpressionUnit) sublexResult {
	var op lexer.Operator
	switch token.Text {
	case "+":
		op = lexer.Addition
		break
	case "-":
		op = lexer.Subtraction
		break
	case "*":
		op = lexer.Multiplication
		break
	case "/":
		op = lexer.Division
		break
	case "^":
		op = lexer.Exponentiation
		break
	case ")":
		ex.exprStack.Pop()
		ex.state = readOperator
		return slValid
	default:
		f := token.Text[0]
		if len(token.Text) == 1 && (f < 'a' || f > 'z') && (f < 'A' || f > 'Z') && (f < '0' || f > '9') {
			return slComplete
		}
		return slError
	}
	*e = &lexer.ExpressionUnit{
		Type:     lexer.OperatorSymbol,
		Operator: op,
	}
	ex.state = readVariable
	return slValid
}

func (ex *expressionReader) Read(token tokenizer.Token) sublexResult {
	t := ex.exprStack.Peek().(*[]lexer.ExpressionUnit)
	var e *lexer.ExpressionUnit = nil
	pushSubExp := false
	switch ex.state {
	case readVariable:
		switch token.Text {
		case "(":
			e = &lexer.ExpressionUnit{
				Type:          lexer.FunctionLiteral,
				Function:      lexer.Parenthesis,
				SubExpression: make([]lexer.ExpressionUnit, 0),
			}
			pushSubExp = true
			break
		case ")":
			ex.exprStack.Pop()
			ex.state = readOperator
			break
		case "sin", "cos", "tan", "csc", "sec", "cot", "asin", "acos", "atan", "acsc", "asec", "acot", "sinh", "cosh", "tanh", "csch", "sech", "coth", "asinh", "acosh", "atanh", "acsch", "asech", "acoth", "exp", "ln":
			var fn lexer.Function
			switch token.Text {
			case "sin":
				fn = lexer.Sine
				break
			case "cos":
				fn = lexer.Cosine
				break
			case "tan":
				fn = lexer.Tangent
				break
			case "csc":
				fn = lexer.Cosecant
				break
			case "sec":
				fn = lexer.Secant
				break
			case "cot":
				fn = lexer.Cotangent
				break
			case "asin":
				fn = lexer.ArcSine
				break
			case "acos":
				fn = lexer.ArcCosine
				break
			case "atan":
				fn = lexer.ArcTangent
				break
			case "acsc":
				fn = lexer.ArcCosecant
				break
			case "asec":
				fn = lexer.ArcSecant
				break
			case "acot":
				fn = lexer.ArcCotangent
				break
			case "sinh":
				fn = lexer.HyperbolicSine
				break
			case "cosh":
				fn = lexer.HyperbolicCosine
				break
			case "tanh":
				fn = lexer.HyperbolicTangent
				break
			case "csch":
				fn = lexer.HyperbolicCosecant
				break
			case "sech":
				fn = lexer.HyperbolicSecant
				break
			case "coth":
				fn = lexer.HyperbolicCotangent
				break
			case "asinh":
				fn = lexer.HyperbolicArcSine
				break
			case "acosh":
				fn = lexer.HyperbolicArcCosine
				break
			case "atanh":
				fn = lexer.HyperbolicArcTangent
				break
			case "acsch":
				fn = lexer.HyperbolicArcCosecant
				break
			case "asech":
				fn = lexer.HyperbolicArcSecant
				break
			case "acoth":
				fn = lexer.HyperbolicArcCotangent
				break
			case "exp":
				fn = lexer.Exponential
				break
			case "ln":
				fn = lexer.Logarithm
				break
			default:
				ex.err(token, errors.MissingCase)
				return slError
			}
			e = &lexer.ExpressionUnit{
				Type:          lexer.FunctionLiteral,
				Function:      fn,
				SubExpression: make([]lexer.ExpressionUnit, 0),
			}
			pushSubExp = true
			ex.state = readOpenParenthesis
			break
		default:
			if isNumber(token.Text) {
				e = &lexer.ExpressionUnit{
					Type: lexer.ExpressionNumber,
					Text: []string{token.Text},
				}
				ex.state = readOperatorOrUnit
			} else if isIdentifier(token.Text) {
				e = &lexer.ExpressionUnit{
					Type: lexer.Variable,
					Text: []string{token.Text},
				}
				ex.state = readVariableDot
			} else {
				ex.err(token, errors.ExpectedExpression, token.Text)
				return slError
			}
			break
		}
		break
	case readOperator:
		switch ex.readOperator(token, &e) {
		case slValid:
			break
		case slComplete:
			return slComplete
		case slError:
			ex.err(token, errors.ExpectedOperator, token.Text)
			return slError
		default:
			ex.err(token, errors.MissingCase)
			return slError
		}
		break
	case readOpenParenthesis:
		if token.Text == "(" {
			ex.state = readVariable
		} else {
			ex.err(token, errors.ExpectedParenthesis, token.Text)
			return slError
		}
		break
	case readVariableDot:
		if token.Text == "." {
			ex.state = readVariableName
		} else {
			switch ex.readOperator(token, &e) {
			case slValid:
				break
			case slComplete:
				return slComplete
			case slError:
				ex.err(token, errors.ExpectedOperator, token.Text)
				return slError
			default:
				ex.err(token, errors.MissingCase)
				return slError
			}
		}
		break
	case readVariableName:
		if isIdentifier(token.Text) {
			u := &(*t)[len(*t)-1]
			u.Text = append(u.Text, token.Text)
			ex.state = readVariableDot
		} else {
			ex.err(token, errors.ExpectedIdentifier, token.Text)
			return slError
		}
		break
	case readOperatorOrUnit:
		valid := false
		switch ex.readOperator(token, &e) {
		case slValid:
			valid = true
			break
		case slComplete:
			return slComplete
		case slError:
			break
		default:
			ex.err(token, errors.MissingCase)
			return slError
		}
		if valid {
			break
		}
		u := &(*t)[len(*t)-1]
		u.Unit = make([]lexer.Unit, 0)
		ex.unit.Reset(&u.Unit)
		ex.state = readUnit
		fallthrough
	case readUnit:
		switch ex.unit.Read(token) {
		case slValid:
			break
		case slComplete:
			ex.state = readOperator
			return ex.Read(token)
		case slError:
			return slError
		default:
			ex.err(token, errors.MissingCase)
			return slError
		}
		break
	default:
		ex.err(token, errors.StateError, ex.state, "expressionState")
		return slError
	}
	if e != nil {
		e.LineNo = token.LineNo
		e.CharNo = token.CharNo
		e.FileName = token.FileName
		*t = append(*t, *e)
		if pushSubExp {
			e = &(*t)[len(*t)-1]
			ex.exprStack.Push(&e.SubExpression)
		}
	}
	return slValid
}
