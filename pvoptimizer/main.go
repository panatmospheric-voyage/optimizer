package main

import (
	"flag"
	"fmt"
	"os"

	"../libpvoptimizer/pipeline"
	resultwriter "../libpvoptimizer/resultwriter/def"
	sourcereader "../libpvoptimizer/sourcereader/def"
)

func main() {
	p := pipeline.CreateDefault()
	outFile := flag.String("o", "", "Output file")
	flag.Parse()
	if *outFile != "" {
		p.ResultWriter.(resultwriter.ResultWriter).SetOutputFile(*outFile)
	}
	switch len(flag.Args()) {
	case 0:
		break
	case 1:
		p.SourceReader.(sourcereader.SourceReader).SetDefaultFile(flag.Args()[0])
		break
	default:
		fmt.Fprintln(os.Stderr, "Error: cannot process multiple files")
		os.Exit(1)
		break
	}
	p.Run()
}
