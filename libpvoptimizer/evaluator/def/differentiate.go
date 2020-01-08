package def

import (
	"strings"

	evaluator ".."
	"../../errors"
	"../../lexer"
	"../../parser"
)

func updateTracking(from parser.Expression, to *parser.Expression) {
	to.LineNo = from.LineNo
	to.CharNo = from.CharNo
	to.FileName = from.FileName
	if to.LHS != nil && to.LHS.LineNo == 0 {
		updateTracking(from, to.LHS)
	}
	if to.RHS != nil && to.RHS.LineNo == 0 {
		updateTracking(from, to.RHS)
	}
}

func differentiate(expr parser.Expression, opt evaluator.OptimizedModel, model parser.Model, respect []string, e errors.IErrorHandler) parser.Expression {
	var res parser.Expression
	null := parser.Expression{
		Type:     parser.Constant,
		Value:    0,
		LineNo:   expr.LineNo,
		CharNo:   expr.CharNo,
		FileName: expr.FileName,
	}
	switch expr.Type {
	case parser.Constant:
		res = parser.Expression{
			Type:  parser.Constant,
			Value: 0,
		}
		break
	case parser.Variable:
		if equals(expr.Name, respect) {
			res = parser.Expression{
				Type:  parser.Constant,
				Value: 1,
			}
			break
		}
		found := false
		for _, p := range opt.Properties {
			if equals(p.Name, expr.Name) {
				res = parser.Expression{
					Type:  parser.Constant,
					Value: 0,
				}
				found = true
				break
			}
		}
		if found {
			break
		}
		for _, p := range model.UniversalProperties {
			if equals(p.Name, expr.Name) {
				return differentiate(p.Definition.RHS, opt, model, respect, e)
			}
		}
		for _, p := range model.Parameters {
			if equals(expr.Name, p.Name) {
				res = expr
				found = true
				break
			}
		}
		if found {
			break
		}
		err(expr, e, errors.UnknownVariable, strings.Join(expr.Name, "."))
		return null
	case parser.Addition, parser.Subtraction:
		lhs := differentiate(*expr.LHS, opt, model, respect, e)
		rhs := differentiate(*expr.RHS, opt, model, respect, e)
		res = parser.Expression{
			Type: expr.Type,
			LHS:  &lhs,
			RHS:  &rhs,
		}
		break
	case parser.Multiplication:
		dl := differentiate(*expr.LHS, opt, model, respect, e)
		dr := differentiate(*expr.RHS, opt, model, respect, e)
		res = parser.Expression{
			Type: parser.Addition,
			LHS: &parser.Expression{
				Type: parser.Multiplication,
				LHS:  &dl,
				RHS:  expr.RHS,
			},
			RHS: &parser.Expression{
				Type: parser.Multiplication,
				LHS:  expr.LHS,
				RHS:  &dr,
			},
		}
		break
	case parser.Division:
		dl := differentiate(*expr.LHS, opt, model, respect, e)
		dr := differentiate(*expr.RHS, opt, model, respect, e)
		res = parser.Expression{
			Type: parser.Division,
			LHS: &parser.Expression{
				Type: parser.Subtraction,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &dl,
					RHS:  expr.RHS,
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  expr.LHS,
					RHS:  &dr,
				},
			},
			RHS: &parser.Expression{
				Type: parser.Multiplication,
				LHS:  expr.RHS,
				RHS:  expr.RHS,
			},
		}
		break
	case parser.Exponentiation:
		dl := differentiate(*expr.LHS, opt, model, respect, e)
		dr := differentiate(*expr.RHS, opt, model, respect, e)
		res = parser.Expression{
			Type: parser.Addition,
			LHS: &parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &dl,
					RHS:  expr.RHS,
				},
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS:  expr.LHS,
					RHS: &parser.Expression{
						Type: parser.Subtraction,
						LHS:  expr.RHS,
						RHS: &parser.Expression{
							Type:  parser.Constant,
							Value: 1,
						},
					},
				},
			},
			RHS: &parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS:  expr.LHS,
					RHS:  expr.RHS,
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &dr,
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Logarithm,
					},
				},
			},
		}
		break
	case parser.Function:
		ins := differentiate(*expr.LHS, opt, model, respect, e)
		switch expr.Function {
		case lexer.Sine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type:     parser.Function,
					LHS:      expr.LHS,
					Function: lexer.Cosine,
				},
			}
			break
		case lexer.Cosine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Sine,
					},
				},
			}
			break
		case lexer.Tangent:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Secant,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Secant,
					},
				},
			}
			break
		case lexer.Cosecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Cotangent,
						LHS:      expr.LHS,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Cosecant,
						LHS:      expr.LHS,
					},
				},
			}
			break
		case lexer.Secant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Tangent,
						LHS:      expr.LHS,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Secant,
						LHS:      expr.LHS,
					},
				},
			}
			break
		case lexer.Cotangent:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Cosecant,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.Cosecant,
					},
				},
			}
			break
		case lexer.ArcSine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type:  parser.Constant,
								Value: 1,
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.LHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.ArcCosine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type:  parser.Constant,
								Value: 1,
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.LHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.ArcTangent:
			res = parser.Expression{
				Type: parser.Division,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Addition,
					LHS: &parser.Expression{
						Type:  parser.Constant,
						Value: 1,
					},
					RHS: &parser.Expression{
						Type: parser.Multiplication,
						LHS:  expr.LHS,
						RHS:  expr.LHS,
					},
				},
			}
			break
		case lexer.ArcCosecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type: parser.Exponentiation,
								LHS:  expr.LHS,
								RHS: &parser.Expression{
									Type:  parser.Constant,
									Value: 4,
								},
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.RHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.ArcSecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type: parser.Exponentiation,
								LHS:  expr.LHS,
								RHS: &parser.Expression{
									Type:  parser.Constant,
									Value: 4,
								},
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.RHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.ArcCotangent:
			res = parser.Expression{
				Type: parser.Division,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Addition,
					LHS: &parser.Expression{
						Type:  parser.Constant,
						Value: 1,
					},
					RHS: &parser.Expression{
						Type: parser.Multiplication,
						LHS:  expr.LHS,
						RHS:  expr.LHS,
					},
				},
			}
			break
		case lexer.HyperbolicSine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type:     parser.Function,
					LHS:      expr.LHS,
					Function: lexer.HyperbolicCosine,
				},
			}
			break
		case lexer.HyperbolicCosine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type:     parser.Function,
					LHS:      expr.LHS,
					Function: lexer.HyperbolicSine,
				},
			}
			break
		case lexer.HyperbolicTangent:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.HyperbolicSecant,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.HyperbolicSecant,
					},
				},
			}
			break
		case lexer.HyperbolicCosecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.HyperbolicCotangent,
						LHS:      expr.LHS,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.HyperbolicCosecant,
						LHS:      expr.LHS,
					},
				},
			}
			break
		case lexer.HyperbolicSecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.HyperbolicTangent,
						LHS:      expr.LHS,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.HyperbolicSecant,
						LHS:      expr.LHS,
					},
				},
			}
			break
		case lexer.HyperbolicCotangent:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.HyperbolicCosecant,
					},
					RHS: &parser.Expression{
						Type:     parser.Function,
						LHS:      expr.LHS,
						Function: lexer.HyperbolicCosecant,
					},
				},
			}
			break
		case lexer.HyperbolicArcSine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Addition,
							LHS: &parser.Expression{
								Type:  parser.Constant,
								Value: 1,
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.LHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.HyperbolicArcCosine:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.LHS,
							},
							RHS: &parser.Expression{
								Type:  parser.Constant,
								Value: 1,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.HyperbolicArcTangent:
			res = parser.Expression{
				Type: parser.Division,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Subtraction,
					LHS: &parser.Expression{
						Type:  parser.Constant,
						Value: 1,
					},
					RHS: &parser.Expression{
						Type: parser.Multiplication,
						LHS:  expr.LHS,
						RHS:  expr.LHS,
					},
				},
			}
			break
		case lexer.HyperbolicArcCosecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Addition,
							LHS: &parser.Expression{
								Type: parser.Exponentiation,
								LHS:  expr.LHS,
								RHS: &parser.Expression{
									Type:  parser.Constant,
									Value: 4,
								},
							},
							RHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.RHS,
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.HyperbolicArcSecant:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS: &parser.Expression{
					Type: parser.Multiplication,
					LHS:  &ins,
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -1,
					},
				},
				RHS: &parser.Expression{
					Type: parser.Exponentiation,
					LHS: &parser.Expression{
						Type:     parser.Function,
						Function: lexer.Parenthesis,
						LHS: &parser.Expression{
							Type: parser.Subtraction,
							LHS: &parser.Expression{
								Type: parser.Multiplication,
								LHS:  expr.LHS,
								RHS:  expr.RHS,
							},
							RHS: &parser.Expression{
								Type: parser.Exponentiation,
								LHS:  expr.LHS,
								RHS: &parser.Expression{
									Type:  parser.Constant,
									Value: 4,
								},
							},
						},
					},
					RHS: &parser.Expression{
						Type:  parser.Constant,
						Value: -0.5,
					},
				},
			}
			break
		case lexer.HyperbolicArcCotangent:
			res = parser.Expression{
				Type: parser.Division,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type: parser.Subtraction,
					LHS: &parser.Expression{
						Type:  parser.Constant,
						Value: 1,
					},
					RHS: &parser.Expression{
						Type: parser.Multiplication,
						LHS:  expr.LHS,
						RHS:  expr.LHS,
					},
				},
			}
			break
		case lexer.Parenthesis:
			res = parser.Expression{
				Type:     parser.Function,
				Function: lexer.Parenthesis,
				LHS:      &ins,
			}
			break
		case lexer.Exponential:
			res = parser.Expression{
				Type: parser.Multiplication,
				LHS:  &ins,
				RHS: &parser.Expression{
					Type:     parser.Function,
					Function: lexer.Exponential,
					LHS:      expr.LHS,
				},
			}
			break
		case lexer.Logarithm:
			res = parser.Expression{
				Type: parser.Division,
				LHS:  &ins,
				RHS:  expr.LHS,
			}
			break
		default:
			err(expr, e, errors.MissingCase)
			return null
		}
		break
	default:
		err(expr, e, errors.MissingCase)
		return null
	}
	updateTracking(expr, &res)
	return res
}
