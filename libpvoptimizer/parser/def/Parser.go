package def

import (
	"fmt"
	"strings"
	"sync"

	parser ".."
	libpvoptimizer "../.."
	"../../errors"
	"../../lexer"
)

// Parser is the default implementation of IParser
type Parser struct {
	handler   errors.IErrorHandler
	evaluator libpvoptimizer.IEvaluator
	model     parser.Model
}

// Init initializes the layer and is called from the pipeline layer
func (ps *Parser) Init(lexer libpvoptimizer.ILexer, evaluator libpvoptimizer.IEvaluator, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ps.handler = e
	ps.evaluator = evaluator
	ps.model = parser.Model{
		Units:               make([]parser.BaseUnit, 0),
		UnitEquivalents:     make([]parser.UnitEquivalence, 0),
		Parameters:          make([]parser.Parameter, 0),
		UniversalProperties: make([]parser.Property, 0),
		Enumerations:        make([]parser.Enumeration, 0),
		Requirements:        make([]parser.Requirement, 0),
		Optimization: parser.Optimization{
			Type: parser.NoOptimize,
		},
	}
	initPatterns()
}

func stream(model *parser.Model, statement lexer.Statement, scope []string, e errors.IErrorHandler) {
	var furthestParse int = 1
	var furthestPatterns []statementPattern = []statementPattern{}
	for _, pat := range patterns {
		if len(statement.Lexemes) == len(pat.parts) {
			valid := true
			for i, part := range pat.parts {
				if part.lType != statement.Lexemes[i].Type || (part.lType == lexer.KeywordLiteral && part.keyword != statement.Lexemes[i].Keyword) {
					if i > furthestParse {
						furthestParse = i
						furthestPatterns = []statementPattern{pat}
					} else if i == furthestParse {
						furthestPatterns = append(furthestPatterns, pat)
					}
					valid = false
					break
				}
			}
			if valid {
				pat.handler(model, statement, scope, e)
				return
			}
		}
	}
	switch len(furthestPatterns) {
	case 0:
		err(statement, e, errors.InvalidStatement, statementPartFromLexeme(statement.Lexemes[0]))
		break
	case 1:
		err(statement, e, errors.NoParserMatch, furthestPatterns[0].parts[furthestParse], statementPartFromLexeme(statement.Lexemes[furthestParse]))
		break
	case 2:
		err(statement, e, errors.NoParserMatch, fmt.Sprintf("%s or %s", furthestPatterns[0].parts[furthestParse], furthestPatterns[1].parts[furthestParse]), statementPartFromLexeme(statement.Lexemes[furthestParse]))
		break
	default:
		exp := make([]string, len(furthestPatterns))
		for i, pat := range furthestPatterns {
			if i == len(exp)-1 {
				exp[i] = fmt.Sprintf("or %s", pat.parts[furthestParse])
			} else {
				exp[i] = pat.parts[furthestParse].String()
			}
		}
		err(statement, e, errors.NoParserMatch, strings.Join(exp, ", "), statementPartFromLexeme(statement.Lexemes[furthestParse]))
		break
	}
}

func err(statement lexer.Statement, e errors.IErrorHandler, code errors.ErrorCode, args ...interface{}) {
	e.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    statement.Lexemes[0].LineNo,
		CharNo:    statement.Lexemes[0].CharNo,
		FileName:  statement.Lexemes[0].FileName,
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

// Stream accepts a statement and processes it.
func (ps *Parser) Stream(statement lexer.Statement) {
	stream(&ps.model, statement, []string{}, ps.handler)
}

// End is called by the lexer once all statements have been streamed into
// the parser.  This is when the model is sent into the evaluator layer.
func (ps *Parser) End() {
	ps.evaluator.Evaluate(ps.model)
}
