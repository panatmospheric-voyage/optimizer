package main

import (
	"flag"
	"fmt"
	"os"

	"../libpvoptimizer/evaluator/mock"
	"../libpvoptimizer/pipeline"
	resultwriter "../libpvoptimizer/resultwriter/def"
	sourcereader "../libpvoptimizer/sourcereader/def"
)

func main() {
	fmt.Println("Hi")
	p := pipeline.CreateDefault()
	p.Evaluator = new(mock.Evaluator)
	outFile := flag.String("o", "", "Output file")
	flag.Parse()
	if *outFile != "" {
		p.ResultWriter.(*resultwriter.ResultWriter).SetOutputFile(*outFile)
	}
	switch len(flag.Args()) {
	case 0:
		break
	case 1:
		p.SourceReader.(*sourcereader.SourceReader).SetDefaultFile(flag.Args()[0])
		break
	default:
		fmt.Fprintln(os.Stderr, "Error: cannot process multiple files")
		os.Exit(1)
		break
	}
	p.Run()
}
