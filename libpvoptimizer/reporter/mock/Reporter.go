package mock

import (
	"fmt"
	"strings"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../evaluator"
	evaluatorMock "../../evaluator/mock"
)

// Reporter is the default implementation of IReporter
type Reporter struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (rp Reporter) Init(evaluator libpvoptimizer.IEvaluator, resultwriter libpvoptimizer.IResultWriter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	rp.handler = e
}

var (
	compOps = []string{"<", "<=", ">", ">=", "==", "!="}
)

// Report generates the report to save
func (rp Reporter) Report(model evaluator.OptimizedModel) {
	fmt.Println("Failed requirements:")
	for _, r := range model.FailedRequirements {
		fmt.Printf("    %s %s ", strings.Join(r.Name, "."), compOps[r.Condition])
		evaluatorMock.PrintExpression(r.Value)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Properties:")
	for _, p := range model.Properties {
		fmt.Printf("    %s = %f %s", strings.Join(p.Name, "."), p.Value, p.Unit)
		if p.Summarize {
			fmt.Print(" [summarize]")
		}
		fmt.Println()
	}
	fmt.Println()
}
