package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

func parseExpressionSingle(expr lexer.ExpressionUnit, model *parser.Model, err errors.IErrorHandler) parser.Expression {
	switch expr.Type {
	case lexer.ExpressionNumber:
		n, e := parser.ParseNumber(expr.Text[0])
		if e != nil {
			err.Handle(errors.Error{
				Arguments: []interface{}{expr.Text[0], e},
				Code:      errors.NumberParseError,
				LineNo:    expr.LineNo,
				CharNo:    expr.CharNo,
				FileName:  expr.FileName,
			})
		}
		if len(expr.Unit) > 0 {
			return parser.Expression{
				Type:  parser.Constant,
				Value: n,
				Unit:  getUnit(model, expr.Unit),
			}
		}
		return parser.Expression{
			Type:  parser.Constant,
			Value: n,
			Unit: parser.Unit{
				Parts: []parser.CompositeUnitPart{},
			},
		}
	case lexer.Variable:
		return parser.Expression{
			Type: parser.Variable,
			Name: expr.Text,
		}
	case lexer.OperatorSymbol:
		err.Handle(errors.Error{
			Arguments: []interface{}{expr.Operator},
			Code:      errors.UnexpectedOperatorError,
			LineNo:    expr.LineNo,
			CharNo:    expr.CharNo,
			FileName:  expr.FileName,
		})
		break
	case lexer.FunctionLiteral:
		e := parseExpression(expr.SubExpression, model, err)
		return parser.Expression{
			Type:     parser.Function,
			LHS:      &e,
			Function: expr.Function,
		}
	default:
		err.Handle(errors.Error{
			Arguments: []interface{}{},
			Code:      errors.MissingCase,
			LineNo:    expr.LineNo,
			CharNo:    expr.CharNo,
			FileName:  expr.FileName,
		})
		break
	}
	return parser.Expression{
		Type:  parser.Constant,
		Value: 0,
		Unit: parser.Unit{
			Parts: []parser.CompositeUnitPart{},
		},
	}
}

func parseExpressionExp(expr []lexer.ExpressionUnit, model *parser.Model, err errors.IErrorHandler) parser.Expression {
	var base *parser.Expression
	it := &base
	var group lexer.ExpressionUnit
	for _, e := range expr {
		if e.Type == lexer.OperatorSymbol && e.Operator == lexer.Exponentiation {
			lhs := parseExpressionSingle(group, model, err)
			*it = &parser.Expression{
				Type: parser.Exponentiation,
				LHS:  &lhs,
			}
			it = &(*it).RHS
		} else {
			group = e
		}
	}
	rhs := parseExpressionSingle(group, model, err)
	*it = &rhs
	return *base
}

func parseExpressionMult(expr []lexer.ExpressionUnit, model *parser.Model, err errors.IErrorHandler) parser.Expression {
	var base *parser.Expression
	it := &base
	group := []lexer.ExpressionUnit{}
	mult := true
	for _, e := range expr {
		if e.Type == lexer.OperatorSymbol {
			if e.Operator == lexer.Multiplication || e.Operator == lexer.Division {
				iMult := mult
				if e.Operator == lexer.Division {
					mult = !mult
				}
				var t parser.ExpressionType
				if mult {
					t = parser.Multiplication
				} else {
					t = parser.Division
				}
				lhs := parseExpressionExp(group, model, err)
				group = []lexer.ExpressionUnit{}
				*it = &parser.Expression{
					Type: t,
					LHS:  &lhs,
				}
				it = &(*it).RHS
				if !iMult {
					mult = true
				}
				continue
			}
		}
		group = append(group, e)
	}
	rhs := parseExpressionExp(group, model, err)
	*it = &rhs
	return *base
}

func parseExpression(expr []lexer.ExpressionUnit, model *parser.Model, err errors.IErrorHandler) parser.Expression {
	var base *parser.Expression
	it := &base
	group := []lexer.ExpressionUnit{}
	plus := true
	for _, e := range expr {
		if e.Type == lexer.OperatorSymbol {
			if e.Operator == lexer.Addition || e.Operator == lexer.Subtraction {
				iPlus := plus
				if e.Operator == lexer.Subtraction {
					plus = !plus
				}
				var t parser.ExpressionType
				if plus {
					t = parser.Addition
				} else {
					t = parser.Subtraction
				}
				lhs := parseExpressionMult(group, model, err)
				group = []lexer.ExpressionUnit{}
				*it = &parser.Expression{
					Type: t,
					LHS:  &lhs,
				}
				it = &(*it).RHS
				if !iPlus {
					plus = true
				}
				continue
			}
		}
		group = append(group, e)
	}
	rhs := parseExpressionMult(group, model, err)
	*it = &rhs
	return *base
}
