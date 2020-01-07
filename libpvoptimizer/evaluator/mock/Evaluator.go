package mock

import (
	"fmt"
	"strings"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../parser"
)

// Evaluator is the default implementation of IEvaluator
type Evaluator struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (ev Evaluator) Init(parser libpvoptimizer.IParser, reporter libpvoptimizer.IReporter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ev.handler = e
}

var (
	compOps = []string{"<", "<=", ">", ">=", "==", "!="}
)

func printUnit(unit parser.Unit) {
	for i, u := range unit.Parts {
		if i != 0 {
			fmt.Print("*")
		}
		fmt.Printf("%s^%d", u.Unit.Name, u.Power)
	}
}

func printExpression(expr parser.Expression) {
	switch expr.Type {
	case parser.Constant:
		fmt.Printf("%f", expr.Value)
		if len(expr.Unit.Parts) > 0 {
			fmt.Print(" ")
			printUnit(expr.Unit)
		}
		break
	case parser.Variable:
		fmt.Print(strings.Join(expr.Name, "."))
		break
	case parser.Addition:
		fmt.Print("(")
		printExpression(*expr.LHS)
		fmt.Print(" + ")
		printExpression(*expr.RHS)
		fmt.Print(")")
		break
	case parser.Subtraction:
		fmt.Print("(")
		printExpression(*expr.LHS)
		fmt.Print(" - ")
		printExpression(*expr.RHS)
		fmt.Print(")")
		break
	case parser.Multiplication:
		fmt.Print("(")
		printExpression(*expr.LHS)
		fmt.Print(" * ")
		printExpression(*expr.RHS)
		fmt.Print(")")
		break
	case parser.Division:
		fmt.Print("(")
		printExpression(*expr.LHS)
		fmt.Print(" / ")
		printExpression(*expr.RHS)
		fmt.Print(")")
		break
	case parser.Exponentiation:
		fmt.Print("(")
		printExpression(*expr.LHS)
		fmt.Print(" ^ ")
		printExpression(*expr.RHS)
		fmt.Print(")")
		break
	case parser.Function:
		fmt.Printf("<func%d>(", expr.Function)
		printExpression(*expr.LHS)
		fmt.Print(")")
		break
	default:
		fmt.Print("?")
		break
	}
}

func printProperties(props []parser.Property, indent string) {
	for _, p := range props {
		fmt.Printf("%s%s", indent, strings.Join(p.Name, "."))
		if p.Summarize {
			fmt.Print(" [summarize]")
		}
		fmt.Print(": ")
		printExpression(p.Definition.LHS)
		fmt.Print(" = ")
		printExpression(p.Definition.RHS)
		fmt.Println()
	}
}

// Evaluate evaluates the model and optimizes it
func (ev Evaluator) Evaluate(model parser.Model) {
	fmt.Print("Base Units: ")
	for i, u := range model.Units {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(u.Name)
	}
	fmt.Println()
	fmt.Println("Unit map:")
	for _, u := range model.UnitEquivalents {
		fmt.Printf("    %s = %f ", u.Unit.Name, u.Factor)
		printUnit(u.EquivalentUnit)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Parameters:")
	for _, p := range model.Parameters {
		var o, c byte
		if p.MinimumInclude {
			o = '['
		} else {
			o = '('
		}
		if p.MaximumInclude {
			c = ']'
		} else {
			c = ')'
		}
		fmt.Printf("    %s %c%f, %f%c ", strings.Join(p.Name, "."), o, p.Minimum, p.Maximum, c)
		printUnit(p.Unit)
		if p.Summarize {
			fmt.Print(" [summarize]")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Properties:")
	printProperties(model.UniversalProperties, "    ")
	fmt.Println()
	fmt.Println("Enumerations:")
	for _, e := range model.Enumerations {
		fmt.Printf("    %s", strings.Join(e.Name, "."))
		if e.Summarize {
			fmt.Print(" [summarize]")
		}
		fmt.Println(":")
		for _, v := range e.Values {
			fmt.Printf("        %s:\n", v.Name)
			printProperties(v.Properties, "           ")
		}
	}
	fmt.Println()
	fmt.Println("Requirements:")
	for _, r := range model.Requirements {
		fmt.Printf("    %s %s ", strings.Join(r.Name, "."), compOps[r.Condition])
		printExpression(r.Value)
		fmt.Println()
	}
	fmt.Println()
	switch model.Type {
	case parser.NoOptimize:
		fmt.Println("No optimization requested.")
		return
	case parser.Minimization:
		fmt.Print("Minimize ")
		break
	case parser.Maximization:
		fmt.Print("Maximize ")
		break
	}
	fmt.Println(strings.Join(model.Variable, "."))
}
