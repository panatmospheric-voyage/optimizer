package def

import (
	"strings"

	evaluator ".."
	"../../errors"
	"../../parser"
)

type gradientState int

type gradientDescent struct {
	handler        errors.IErrorHandler
	opt            evaluator.OptimizedModel
	parameters     []parser.Parameter
	gradientVector []parser.Expression
	gradientScale  []parser.Number
	step           parser.Number
	model          parser.Model
	lastEval       parser.Number
	expr           parser.Expression
}

func (gd *gradientDescent) err(code errors.ErrorCode, args ...interface{}) {
	gd.handler.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    gd.model.Optimization.LineNo,
		CharNo:    gd.model.Optimization.CharNo,
		FileName:  gd.model.Optimization.FileName,
	})
}

func (gd *gradientDescent) init(opt evaluator.OptimizedModel, model parser.Model, e errors.IErrorHandler, iter int) {
	gd.handler = e
	gd.opt = opt
	gd.parameters = model.Parameters
	gd.gradientVector = make([]parser.Expression, len(gd.parameters))
	gd.gradientScale = make([]parser.Number, len(gd.parameters))
	found := false
	for _, p := range model.UniversalProperties {
		if equals(p.Name, model.Optimization.Variable) {
			gd.expr = p.Definition.RHS
			found = true
			break
		}
	}
	if !found {
		gd.err(errors.UnknownVariable, strings.Join(model.Optimization.Variable, "."))
		return
	}
	var state uint64 = (0xFA1150FF5CAFF01D ^ 2*(uint64(len(model.Parameters))*uint64(iter)+uint64(model.Optimization.Seed))) | 1
	gd.opt.Properties = append(make([]evaluator.Property, len(gd.parameters)), gd.opt.Properties...)
	for i, param := range gd.parameters {
		gd.gradientVector[i] = differentiate(gd.expr, opt, model, param.Name, e)
		diff := param.Maximum - param.Minimum
		gd.gradientScale[i] = diff
		x := state
		count := x >> 61
		state = x * 0xD1AB011CA1D00D1E
		x ^= x >> 22
		r := uint32(x >> (22 + count))
		gd.opt.Properties[i] = evaluator.Property{
			Name:      param.Name,
			Value:     param.Minimum + diff*parser.Number(r)/parser.Number(0xFFFFFFFF),
			Unit:      param.Unit,
			Summarize: param.Summarize,
		}
	}
	gd.step = 0.5
	gd.model = model
	v, _ := evaluate(gd.expr, gd.opt, gd.model, gd.handler, false)
	if v == nil {
		gd.err(errors.CannotEvaluateObjective)
		return
	}
	gd.lastEval = *v
}

func (gd *gradientDescent) run() (parser.Number, parser.Number) {
	gradient := make([]parser.Number, len(gd.gradientVector))
	var mag2 parser.Number = 0
	for i, vect := range gd.gradientVector {
		r, _ := evaluate(vect, gd.opt, gd.model, gd.handler, true)
		if r == nil {
			err(vect, gd.handler, errors.CannotEvaluateDerivative)
			return 0, 0
		}
		gradient[i] = *r
		mag2 += *r * *r
	}
	f := 1 / mag2.Pow(0.5)
	for i := range gradient {
		gradient[i] *= f * gd.step * gd.gradientScale[i]
	}
	var flip bool
	switch gd.model.Optimization.Type {
	case parser.Minimization:
		flip = true
		break
	case parser.Maximization:
		flip = false
		break
	case parser.Zero:
		flip = gd.lastEval > 0
		break
	default:
		gd.err(errors.MissingCase)
		return 0, 0
	}
	if flip {
		for i := range gradient {
			gradient[i] *= -1
		}
	}
	for i, g := range gradient {
		n := &gd.opt.Properties[i].Value
		p := gd.model.Parameters[i]
		*n += g
		if *n <= p.Minimum && !p.MinimumInclude {
			*n = p.Minimum + gd.model.Optimization.Minimum
		} else if *n < p.Minimum {
			*n = p.Minimum
		}
		if *n >= p.Maximum && !p.MaximumInclude {
			*n = p.Maximum + gd.model.Optimization.Minimum
		} else if *n > p.Maximum {
			*n = p.Maximum
		}
	}
	v, _ := evaluate(gd.expr, gd.opt, gd.model, gd.handler, false)
	if v == nil {
		gd.err(errors.CannotEvaluateObjective)
		return 0, 0
	}
	diff := parser.Abs(*v - gd.lastEval)
	gd.lastEval = *v
	gd.step /= 2
	return diff, *v
}

func (gd *gradientDescent) getModel() evaluator.OptimizedModel {
	evaluateConstants(&gd.opt, &gd.model, gd.handler)
	if len(gd.model.UniversalProperties) > 0 {
		gd.err(errors.UnableToSolveFully)
	}
	gd.opt.FailedRequirements = gd.model.Requirements
	return gd.opt
}
