package def

import (
	"fmt"
	"strings"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../evaluator"
	"../../parser"
)

// Reporter is the default implementation of IReporter
type Reporter struct {
	handler      errors.IErrorHandler
	resultwriter libpvoptimizer.IResultWriter
}

// Init initializes the layer and is called from the pipeline layer
func (rp *Reporter) Init(evaluator libpvoptimizer.IEvaluator, resultwriter libpvoptimizer.IResultWriter, e errors.IErrorHandler, wg *sync.WaitGroup) {
	rp.handler = e
	rp.resultwriter = resultwriter
}

type node struct {
	name     string
	value    parser.Number
	unit     parser.Unit
	children []node
}

func printFullReport(b *strings.Builder, node node, indent string) {
	if len(node.children) > 0 {
		fmt.Fprintf(b, "%s<%s>\n", indent, node.name)
		nIndent := fmt.Sprintf("    %s", indent)
		for _, n := range node.children {
			printFullReport(b, n, nIndent)
		}
		fmt.Fprintf(b, "%s</%s>\n", indent, node.name)
	} else if len(node.unit.Parts) > 0 {
		fmt.Fprintf(b, "%s<%s units=\"%s\">%s</%s>\n", indent, node.name, node.unit, node.value, node.name)
	} else {
		fmt.Fprintf(b, "%s<%s>%s</%s>\n", indent, node.name, node.value, node.name)
	}
}

// Report generates the report to save
func (rp Reporter) Report(model evaluator.OptimizedModel) {
	var b strings.Builder
	fmt.Fprintln(&b, "<?xml version=\"1.0\"?>")
	fmt.Fprintln(&b, "<optimization>")
	fmt.Fprintln(&b, "    <summary>")
	nodes := []node{}
	for _, p := range model.Properties {
		if p.Summarize {
			fmt.Fprintf(&b, "        %s = %s %s\n", strings.Join(p.Name, "."), p.Value, p.Unit)
		}
		var n *node
		l := &nodes
		for _, v := range p.Name {
			found := false
			for i, w := range *l {
				if w.name == v {
					n = &(*l)[i]
					found = true
					break
				}
			}
			if !found {
				i := len(*l)
				*l = append(*l, node{
					name:     v,
					children: []node{},
				})
				n = &(*l)[i]
			}
			l = &n.children
		}
		n.value = p.Value
		n.unit = p.Unit
	}
	fmt.Fprintln(&b, "    </summary>")
	if len(model.FailedRequirements) > 0 {
		fmt.Fprintln(&b, "    <failedreqs>")
		for _, r := range model.FailedRequirements {
			fmt.Fprintf(&b, "        <failedreq variable=\"%s\" line=\"%d\" char=\"%d\" file=\"%s\" />\n", strings.Join(r.Name, "."), r.Value.LineNo, r.Value.CharNo, r.Value.FileName)
		}
		fmt.Fprintln(&b, "    </failedreqs>")
	}
	fmt.Fprintln(&b, "    <full>")
	for _, n := range nodes {
		printFullReport(&b, n, "        ")
	}
	fmt.Fprintln(&b, "    </full>")
	fmt.Fprintln(&b, "</optimization>")
	rp.resultwriter.Save(b.String())
}
