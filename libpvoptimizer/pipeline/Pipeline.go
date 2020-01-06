package pipeline

import ".."

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
}
