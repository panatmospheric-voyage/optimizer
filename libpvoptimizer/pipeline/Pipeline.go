package pipeline

import (
	"sync"

	libpvoptimizer ".."
	"../errors"
	errorhandler "../errors/def"
	evaluator "../evaluator/def"
	lexer "../lexer/def"
	parser "../parser/def"
	reporter "../reporter/def"
	resultwriter "../resultwriter/def"
	sourcereader "../sourcereader/def"
	tokenizer "../tokenizer/def"
)

// Pipeline represents the optimization pipeline
type Pipeline struct {
	// SourceReader layer
	SourceReader libpvoptimizer.ISourceReader
	// Tokenizer layer
	Tokenizer libpvoptimizer.ITokenizer
	// Lexer layer
	Lexer libpvoptimizer.ILexer
	// Parser layer
	Parser libpvoptimizer.IParser
	// Evalutor layer
	Evaluator libpvoptimizer.IEvaluator
	// Reporter layer
	Reporter libpvoptimizer.IReporter
	// ResultWriter layer
	ResultWriter libpvoptimizer.IResultWriter
	// ErrorHandler layer
	ErrorHandler errors.IErrorHandler
}

// Run runs the pipeline
func (p Pipeline) Run() {
	var wg sync.WaitGroup
	p.SourceReader.Init(p.Tokenizer, p.ErrorHandler, &wg)
	p.Tokenizer.Init(p.SourceReader, p.Lexer, p.ErrorHandler, &wg)
	p.Lexer.Init(p.Tokenizer, p.Parser, p.ErrorHandler, &wg)
	p.Parser.Init(p.Lexer, p.Evaluator, p.ErrorHandler, &wg)
	p.Evaluator.Init(p.Parser, p.Reporter, p.ErrorHandler, &wg)
	p.Reporter.Init(p.Evaluator, p.ResultWriter, p.ErrorHandler, &wg)
	p.ResultWriter.Init(p.Reporter, p.ErrorHandler, &wg)
	wg.Wait()
}

// CreateDefault creates a pipeline with the default implementations
func CreateDefault() Pipeline {
	return Pipeline{
		SourceReader: new(sourcereader.SourceReader),
		Tokenizer:    new(tokenizer.Tokenizer),
		Lexer:        new(lexer.Lexer),
		Parser:       new(parser.Parser),
		Evaluator:    new(evaluator.Evaluator),
		Reporter:     new(reporter.Reporter),
		ResultWriter: new(resultwriter.ResultWriter),
		ErrorHandler: new(errorhandler.ErrorHandler),
	}
}
