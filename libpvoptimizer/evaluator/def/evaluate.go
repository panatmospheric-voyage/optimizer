package def

import (
	evaluator ".."
	"../../errors"
	"../../lexer"
	"../../parser"
)

func evaluate(expr parser.Expression, opt evaluator.OptimizedModel, model parser.Model, e errors.IErrorHandler, ignoreUnits bool) (*parser.Number, *parser.Unit) {
	switch expr.Type {
	case parser.Constant:
		return &expr.Value, &expr.Unit
	case parser.Variable:
		for _, p := range opt.Properties {
			if equals(p.Name, expr.Name) {
				return &p.Value, &p.Unit
			}
		}
		for _, p := range model.UniversalProperties {
			if equals(p.Name, expr.Name) {
				return evaluate(p.Definition.RHS, opt, model, e, ignoreUnits)
			}
		}
		break
	case parser.Addition, parser.Subtraction, parser.Multiplication, parser.Division, parser.Exponentiation:
		lhs, lunit := evaluate(*expr.LHS, opt, model, e, ignoreUnits)
		rhs, runit := evaluate(*expr.RHS, opt, model, e, ignoreUnits)
		if lhs == nil || rhs == nil {
			return nil, nil
		}
		var res parser.Number
		var u parser.Unit
		switch expr.Type {
		case parser.Addition:
			res = *lhs + *rhs
			if !unitEquals(*lunit, *runit) && *lhs != 0 && *rhs != 0 && !ignoreUnits {
				err(expr, e, errors.AddUnitMismatch, *lunit, *runit)
				return nil, nil
			}
			if *lhs == 0 {
				u = *runit
			} else {
				u = *lunit
			}
			break
		case parser.Subtraction:
			res = *lhs - *rhs
			if !unitEquals(*lunit, *runit) && *lhs != 0 && *rhs != 0 && !ignoreUnits {
				err(expr, e, errors.SubtractUnitMismatch, *lunit, *runit)
				return nil, nil
			}
			if *lhs == 0 {
				u = *runit
			} else {
				u = *lunit
			}
			break
		case parser.Multiplication:
			res = *lhs * *rhs
			u = unitMultiply(*lunit, *runit)
			break
		case parser.Division:
			res = *lhs / *rhs
			u = unitDivide(*lunit, *runit)
			break
		case parser.Exponentiation:
			res = lhs.Pow(*rhs)
			if len(runit.Parts) != 0 && !ignoreUnits {
				err(expr, e, errors.ShouldBeUnitless, *runit)
				return nil, nil
			}
			u = parser.Unit{
				Parts: make([]parser.CompositeUnitPart, len(lunit.Parts)),
			}
			copy(u.Parts, lunit.Parts)
			for i := range u.Parts {
				u.Parts[i].Power *= *rhs
			}
			break
		default:
			err(expr, e, errors.MissingCase)
			return nil, nil
		}
		n := 0
		for _, p := range u.Parts {
			if p.Power != 0 {
				u.Parts[n] = p
				n++
			}
		}
		u.Parts = u.Parts[:n]
		return &res, &u
	case parser.Function:
		var res parser.Number
		px, u := evaluate(*expr.LHS, opt, model, e, ignoreUnits)
		if px == nil {
			return nil, nil
		}
		if expr.Function != lexer.Parenthesis && len(u.Parts) != 0 && !ignoreUnits {
			err(expr, e, errors.ShouldBeUnitless, *u)
			return nil, nil
		}
		x := *px
		switch expr.Function {
		case lexer.Sine:
			res = parser.Sine(x)
			break
		case lexer.Cosine:
			res = parser.Cosine(x)
			break
		case lexer.Tangent:
			res = parser.Tangent(x)
			break
		case lexer.Cosecant:
			res = 1 / parser.Sine(x)
			break
		case lexer.Secant:
			res = 1 / parser.Cosine(x)
			break
		case lexer.Cotangent:
			res = 1 / parser.Tangent(x)
			break
		case lexer.ArcSine:
			res = parser.ArcSine(x)
			break
		case lexer.ArcCosine:
			res = parser.ArcCosine(x)
			break
		case lexer.ArcTangent:
			res = parser.ArcTangent(x)
			break
		case lexer.ArcCosecant:
			res = parser.ArcSine(1 / x)
			break
		case lexer.ArcSecant:
			res = parser.ArcCosine(1 / x)
			break
		case lexer.ArcCotangent:
			res = parser.ArcTangent(1 / x)
			break
		case lexer.HyperbolicSine:
			res = parser.HyperbolicSine(x)
			break
		case lexer.HyperbolicCosine:
			res = parser.HyperbolicCosine(x)
			break
		case lexer.HyperbolicTangent:
			res = parser.HyperbolicTangent(x)
			break
		case lexer.HyperbolicCosecant:
			res = 1 / parser.HyperbolicSine(x)
			break
		case lexer.HyperbolicSecant:
			res = 1 / parser.HyperbolicCosine(x)
			break
		case lexer.HyperbolicCotangent:
			res = 1 / parser.HyperbolicTangent(x)
			break
		case lexer.HyperbolicArcSine:
			res = parser.HyperbolicArcSine(x)
			break
		case lexer.HyperbolicArcCosine:
			res = parser.HyperbolicArcCosine(x)
			break
		case lexer.HyperbolicArcTangent:
			res = parser.HyperbolicArcTangent(x)
			break
		case lexer.HyperbolicArcCosecant:
			res = parser.HyperbolicArcSine(1 / x)
			break
		case lexer.HyperbolicArcSecant:
			res = parser.HyperbolicArcCosine(1 / x)
			break
		case lexer.HyperbolicArcCotangent:
			res = parser.HyperbolicArcTangent(1 / x)
			break
		case lexer.Parenthesis:
			res = x
			break
		case lexer.Exponential:
			res = parser.Exponential(x)
			break
		case lexer.Logarithm:
			res = parser.Logarithm(x)
			break
		default:
			err(expr, e, errors.MissingCase)
			return nil, nil
		}
		return &res, u
	default:
		err(expr, e, errors.MissingCase)
		break
	}
	return nil, nil
}
