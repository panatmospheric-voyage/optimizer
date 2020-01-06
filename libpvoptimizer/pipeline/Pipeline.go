package pipeline

import (
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

// Run runs the pipeline (asynchronously)
func (p Pipeline) Run() {
	p.SourceReader.Init(p.Tokenizer, p.ErrorHandler)
	p.Tokenizer.Init(p.SourceReader, p.Lexer, p.ErrorHandler)
	p.Lexer.Init(p.Tokenizer, p.Parser, p.ErrorHandler)
	p.Parser.Init(p.Lexer, p.Evaluator, p.ErrorHandler)
	p.Evaluator.Init(p.Parser, p.Reporter, p.ErrorHandler)
	p.Reporter.Init(p.Evaluator, p.ResultWriter, p.ErrorHandler)
	p.ResultWriter.Init(p.Reporter, p.ErrorHandler)
}

// CreateDefault creates a pipeline with the default implementations
func CreateDefault() Pipeline {
	return Pipeline{
		SourceReader: sourcereader.SourceReader{},
		Tokenizer:    tokenizer.Tokenizer{},
		Lexer:        lexer.Lexer{},
		Parser:       parser.Parser{},
		Evaluator:    evaluator.Evaluator{},
		Reporter:     reporter.Reporter{},
		ResultWriter: resultwriter.ResultWriter{},
		ErrorHandler: errorhandler.ErrorHandler{},
	}
}
