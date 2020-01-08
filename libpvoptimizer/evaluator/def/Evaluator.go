package def

import (
	"strings"
	"sync"

	evaluator ".."
	libpvoptimizer "../.."
	"../../errors"
	"../../parser"
)

// Evaluator is the default implementation of IEvaluator
type Evaluator struct {
	handler  errors.IErrorHandler
	reporter libpvoptimizer.IReporter
}

func err(expr parser.Expression, e errors.IErrorHandler, code errors.ErrorCode, args ...interface{}) {
	e.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    expr.LineNo,
		CharNo:    expr.CharNo,
		FileName:  expr.FileName,
	})
}

func equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Init initializes the layer and is called from the pipeline layer
func (ev *Evaluator) Init(parser libpvoptimizer.IParser, reporter libpvoptimizer.IReporter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ev.handler = e
	ev.reporter = reporter
}

func evaluateConstants(opt *evaluator.OptimizedModel, model *parser.Model, e errors.IErrorHandler) {
	for evaluated := true; evaluated; {
		evaluated = false
		i := 0
		for _, p := range model.UniversalProperties {
			n, u := evaluate(p.Definition.RHS, *opt, *model, e, false)
			if n != nil {
				evaluated = true
				opt.Properties = append(opt.Properties, evaluator.Property{
					Name:      p.Name,
					Value:     *n,
					Unit:      *u,
					Summarize: p.Summarize,
				})
			} else {
				model.UniversalProperties[i] = p
				i++
			}
		}
		model.UniversalProperties = model.UniversalProperties[:i]
	}
}

// Evaluate evaluates the model and optimizes it
func (ev *Evaluator) Evaluate(model parser.Model) {
	n := 0
	for _, p := range model.UniversalProperties {
		if !solve(&p.Definition, p.Name, ev.handler) {
			err(p.Definition.LHS, ev.handler, errors.CannotSolveEquation, strings.Join(p.Name, "."))
		} else {
			model.UniversalProperties[n] = p
			n++
		}
	}
	model.UniversalProperties = model.UniversalProperties[:n]
	opt := evaluator.OptimizedModel{}
	evaluateConstants(&opt, &model, ev.handler)
	var bestValue parser.Number
	var bestModel evaluator.OptimizedModel
	var gd gradientDescent
	iters := model.Optimization.Iterations
	if iters <= 0 {
		iters = 3 * len(model.Parameters)
	}
	for i := 0; i < iters; i++ {
		gd.init(opt, model, ev.handler, i)
		acc, val := gd.run()
		for acc > model.Optimization.Accuracy {
			acc, val = gd.run()
		}
		better := false
		switch model.Optimization.Type {
		case parser.Minimization:
			better = val < bestValue
			break
		case parser.Maximization:
			better = val > bestValue
			break
		case parser.Zero:
			better = parser.Abs(val) < parser.Abs(bestValue)
			break
		default:
			ev.handler.Handle(errors.Error{
				Arguments: []interface{}{},
				Code:      errors.MissingCase,
				LineNo:    model.Optimization.LineNo,
				CharNo:    model.Optimization.CharNo,
				FileName:  model.Optimization.FileName,
			})
			break
		}
		if i == 0 || better {
			bestValue = val
			bestModel = gd.getModel()
		}
	}
	ev.reporter.Report(bestModel)
}
