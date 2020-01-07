package def

import (
	parser ".."
	"../../errors"
	"../../lexer"
)

type statementPart struct {
	lType   lexer.LexemeType
	keyword lexer.Keyword
}

func (s statementPart) String() string {
	if s.lType == lexer.KeywordLiteral {
		return s.keyword.String()
	}
	return s.lType.String()
}

func statementPartFromLexeme(l lexer.Lexeme) statementPart {
	return statementPart{
		lType:   l.Type,
		keyword: l.Keyword,
	}
}

type statementPattern struct {
	parts   []statementPart
	handler func(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler)
}

var _patterns []statementPattern = []statementPattern{
	// Unit equivalence
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.UnitKeyword},
			{lexer.UnitLiteral, 0},
			{lexer.KeywordLiteral, lexer.EqualsKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.UnitLiteral, 0},
		},
		handler: parseUnitEquivalence,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.UnitKeyword},
			{lexer.UnitLiteral, 0},
			{lexer.KeywordLiteral, lexer.EqualsKeyword},
			{lexer.UnitLiteral, 0},
		},
		handler: parseUnitExactEquivalence,
	},
	// Properties
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.PropertyKeyword},
			{lexer.Identifier, 0},
		},
		handler: parsePropertyPrototype,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.PropertyKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.EqualsKeyword},
			{lexer.Expression, 0},
		},
		handler: parsePropertyDefinition,
	},
	{
		parts: []statementPart{
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.EqualsKeyword},
			{lexer.Expression, 0},
		},
		handler: parsePropertyAssignment,
	},
	// Parameters
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.InclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.InclusiveClose},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.InclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveClose},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.InclusiveClose},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveClose},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.InclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.InclusiveClose},
			{lexer.UnitLiteral, 0},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.InclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveClose},
			{lexer.UnitLiteral, 0},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.InclusiveClose},
			{lexer.UnitLiteral, 0},
		},
		handler: parseParameter,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ParameterKeyword},
			{lexer.Identifier, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveOpen},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.CommaKeyword},
			{lexer.NumberLiteral, 0},
			{lexer.KeywordLiteral, lexer.ExclusiveClose},
			{lexer.UnitLiteral, 0},
		},
		handler: parseParameter,
	},
	// Unaries
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.SummarizeKeyword},
			{lexer.Expression, 0},
		},
		handler: parseSummarize,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.MinimizeKeyword},
			{lexer.Expression, 0},
		},
		handler: parseMinimize,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.MaximizeKeyword},
			{lexer.Expression, 0},
		},
		handler: parseMaximize,
	},
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.ZeroKeyword},
			{lexer.Expression, 0},
		},
		handler: parseZero,
	},
	// Requirement
	{
		parts: []statementPart{
			{lexer.KeywordLiteral, lexer.RequireKeyword},
			{lexer.Expression, 0},
			{lexer.Conditional, 0},
			{lexer.Expression, 0},
		},
		handler: parseRequirement,
	},
	// Blocks
	{
		parts: []statementPart{
			{lexer.GroupBlock, 0},
		},
		handler: parseBlock,
	},
	// If/else
	{
		parts: []statementPart{
			{lexer.Switch, 0},
		},
		handler: parseIfElse,
	},
}
var patterns []statementPattern

func initPatterns() {
	// Go has a weird limitation on this
	patterns = _patterns
}
