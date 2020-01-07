package def

import (
	"sync"

	lexer ".."
	libpvoptimizer "../.."
	"../../errors"
	"../../sourcereader"
	"../../tokenizer"
	"github.com/golang-collections/collections/stack"
)

type scope struct {
	lexemes    []*lexer.Lexeme
	statements []lexer.Statement
	wasIf      bool
}

// Lexer is the default implementation of ILexer
type Lexer struct {
	handler   errors.IErrorHandler
	tokenizer libpvoptimizer.ITokenizer
	parser    libpvoptimizer.IParser
	scopes    *stack.Stack
	state     lexerState
	unit      unitReader
	expr      expressionReader
	tmp       byte
}

func (lx *Lexer) startScope() {
	lx.scopes.Push(&scope{
		lexemes:    make([]*lexer.Lexeme, 0),
		statements: make([]lexer.Statement, 0),
		wasIf:      false,
	})
	lx.state = statementStart
}

// Init initializes the layer and is called from the pipeline layer
func (lx *Lexer) Init(tokenizer libpvoptimizer.ITokenizer, parser libpvoptimizer.IParser, e errors.IErrorHandler, wg *sync.WaitGroup) {
	lx.handler = e
	lx.tokenizer = tokenizer
	lx.parser = parser
	lx.scopes = stack.New()
	lx.startScope()
	lx.unit.Init(e)
	lx.expr.Init(e)
	tokenizer.ReadFile("", 0, sourcereader.DefaultSource)
}

func (lx *Lexer) err(token tokenizer.Token, code errors.ErrorCode, args ...interface{}) {
	lx.handler.Handle(errors.Error{
		Arguments: args,
		Code:      code,
		LineNo:    token.LineNo,
		CharNo:    token.CharNo,
		FileName:  token.FileName,
	})
}

func (lx *Lexer) finishStatement() {
	lx.state = statementStart
	sc := lx.scopes.Peek().(*scope)
	st := lexer.Statement{
		Lexemes: make([]lexer.Lexeme, len(sc.lexemes)),
	}
	for i, l := range sc.lexemes {
		st.Lexemes[i] = *l
	}
	sc.lexemes = make([]*lexer.Lexeme, 0)
	if lx.scopes.Len() == 1 {
		lx.parser.Stream(st)
	} else {
		sc.statements = append(sc.statements, st)
	}
}

// Stream accepts a token and processes it.  If it is the end of a
// statement, it then streams it to the parser layer.
func (lx *Lexer) Stream(token tokenizer.Token, id int) {
	for repeat := true; repeat; {
		repeat = false
		sc := lx.scopes.Peek().(*scope)
		var lexeme *lexer.Lexeme = nil
		switch lx.state {
		case statementStart:
			if sc.wasIf {
				if token.Text == "else" {
					lx.state = readElse
					break
				} else {
					sc.wasIf = false
				}
			}
			switch token.Text {
			case ";":
				lx.err(token, errors.EmptyStatement)
				break
			case "unit":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.UnitKeyword,
				}
				lx.state = readUnitLHS
				break
			case "property":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.PropertyKeyword,
				}
				lx.state = readPropertyName
				break
			case "assembly", "enum", "value":
				lexeme = &lexer.Lexeme{
					Type: lexer.GroupBlock,
				}
				switch token.Text {
				case "assembly":
					lexeme.Keyword = lexer.AssemblyKeyword
					break
				case "enum":
					lexeme.Keyword = lexer.EnumKeyword
					break
				case "value":
					lexeme.Keyword = lexer.ValueKeyword
					break
				default:
					lx.err(token, errors.MissingCase)
					break
				}
				lx.state = readBlockName
				break
			case "parameter":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.ParameterKeyword,
				}
				lx.state = readParameterName
				break
			case "summarize", "minimize", "maximize":
				lexeme = &lexer.Lexeme{
					Type: lexer.KeywordLiteral,
				}
				switch token.Text {
				case "summarize":
					lexeme.Keyword = lexer.SummarizeKeyword
					break
				case "minimize":
					lexeme.Keyword = lexer.MinimizeKeyword
					break
				case "maximize":
					lexeme.Keyword = lexer.MaximizeKeyword
					break
				default:
					lx.err(token, errors.MissingCase)
					break
				}
				lx.state = readUnaryStatementName
				break
			case "require":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.RequireKeyword,
				}
				lx.state = readRequirementName
				break
			case "if":
				lexeme = &lexer.Lexeme{
					Type: lexer.Switch,
					SwitchBlocks: []lexer.SwitchBlock{{
						LineNo:   token.LineNo,
						CharNo:   token.CharNo,
						FileName: token.FileName,
					}},
				}
				lx.expr.Reset(&lexeme.SwitchBlocks[0].LHS)
				lx.state = readIfLHS
				sc.wasIf = true
				break
			case "else":
				lx.err(token, errors.UnexpectedElse)
				break
			case "}":
				if len(sc.lexemes) != 0 {
					lx.err(token, errors.StateError, len(sc.lexemes), "len(sc.lexemes)")
				}
				if lx.scopes.Len() > 1 {
					lx.scopes.Pop()
					sc2 := lx.scopes.Peek().(*scope)
					lm := sc2.lexemes[len(sc2.lexemes)-1]
					if sc2.wasIf {
						lm.SwitchBlocks[len(lm.SwitchBlocks)-1].Statements = sc.statements
					} else {
						lm.Statements = sc.statements
						lx.finishStatement()
					}
				} else {
					lx.err(token, errors.UnexpectedEndBlock)
				}
				break
			default:
				if isIdentifier(token.Text) {
					lexeme = &lexer.Lexeme{
						Type: lexer.Identifier,
						Name: token.Text,
					}
					lx.state = readPropertyEquals
				} else {
					lx.err(token, errors.ExpectedStatement, token.Text)
				}
				break
			}
			break
		case readUnitLHS:
			if isIdentifier(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.UnitLiteral,
					Unit: []lexer.Unit{{
						Name:     token.Text,
						Power:    1,
						LineNo:   token.LineNo,
						CharNo:   token.CharNo,
						FileName: token.FileName,
					}},
				}
				lx.state = readUnitEquals
			} else {
				lx.err(token, errors.ExpectedUnit, token.Text)
				lx.finishStatement()
			}
			break
		case readUnitEquals:
			if token.Text == "=" {
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.EqualsKeyword,
				}
				lx.state = readUnitRHS
			} else {
				lx.err(token, errors.ExpectedEquals, token.Text)
				lx.finishStatement()
			}
			break
		case readUnitRHS:
			if isNumber(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.NumberLiteral,
					Name: token.Text,
				}
				lx.state = readUnitEnd
				break
			} else if !isIdentifier(token.Text) {
				lx.err(token, errors.ExpectedNumberOrUnit, token.Text)
				lx.finishStatement()
				break
			}
			lx.state = readUnitEnd
			fallthrough
		case readUnitEnd:
			l := sc.lexemes[len(sc.lexemes)-1]
			if l.Type != lexer.UnitLiteral {
				lexeme = &lexer.Lexeme{
					Type: lexer.UnitLiteral,
				}
				lx.unit.Reset(&lexeme.Unit)
			}
			switch lx.unit.Read(token) {
			case slValid:
				break
			case slComplete:
				if token.Text != ";" {
					lx.err(token, errors.ExpectedEndOfStatement, token.Text)
				}
				lx.finishStatement()
				break
			case slError:
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.MissingCase)
				lx.finishStatement()
				break
			}
			break
		case readEndOfStatement:
			if token.Text != ";" {
				lx.err(token, errors.ExpectedEndOfStatement, token.Text)
			}
			lx.finishStatement()
			break
		case readPropertyName:
			if isIdentifier(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.Identifier,
					Name: token.Text,
				}
				lx.state = readPropertyEquals
			} else {
				lx.err(token, errors.ExpectedIdentifier, token.Text)
				lx.finishStatement()
			}
			break
		case readPropertyEquals:
			switch token.Text {
			case "=":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.EqualsKeyword,
				}
				lx.state = readExpression
				break
			case "<-":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.AssignArrow,
				}
				lx.state = readPropertyLHS
				break
			case ";":
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.ExpectedEquals, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readExpression:
			l := sc.lexemes[len(sc.lexemes)-1]
			if l.Type != lexer.Expression {
				lexeme = &lexer.Lexeme{
					Type:       lexer.Expression,
					Expression: make([]lexer.ExpressionUnit, 0),
				}
				lx.expr.Reset(&lexeme.Expression)
			}
			switch lx.expr.Read(token) {
			case slValid:
				break
			case slComplete:
				if token.Text != ";" {
					lx.err(token, errors.ExpectedEndOfStatement, token.Text)
				}
				lx.finishStatement()
				break
			case slError:
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.MissingCase)
				lx.finishStatement()
				break
			}
			break
		case readBlockName:
			if isIdentifier(token.Text) {
				sc.lexemes[len(sc.lexemes)-1].Name = token.Text
				lx.state = readBlockStart
			} else {
				lx.err(token, errors.ExpectedIdentifier, token.Text)
				lx.finishStatement()
			}
			break
		case readBlockStart:
			if token.Text == "{" {
				lx.startScope()
			} else {
				lx.err(token, errors.ExpectedBlockStart, token.Text)
				lx.finishStatement()
			}
			break
		case readParameterName:
			if isIdentifier(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.Identifier,
					Name: token.Text,
				}
				lx.state = readParameterOpen
			} else {
				lx.err(token, errors.ExpectedIdentifier, token.Text)
				lx.finishStatement()
			}
			break
		case readParameterOpen:
			switch token.Text {
			case "(":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.ExclusiveOpen,
				}
				lx.state = readParameterLower
				break
			case "[":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.InclusiveOpen,
				}
				lx.state = readParameterLower
				break
			default:
				lx.err(token, errors.ExpectedRangeOpen, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readParameterLower:
			if isNumber(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.NumberLiteral,
					Name: token.Text,
				}
				lx.state = readParameterDelim
			} else {
				lx.err(token, errors.ExpectedNumber, token.Text)
				lx.finishStatement()
			}
			break
		case readParameterDelim:
			if token.Text == "," {
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.CommaKeyword,
				}
				lx.state = readParameterUpper
			} else {
				lx.err(token, errors.ExpectedDelimiter, token.Text)
				lx.finishStatement()
			}
			break
		case readParameterUpper:
			if isNumber(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.NumberLiteral,
					Name: token.Text,
				}
				lx.state = readParameterClose
			} else {
				lx.err(token, errors.ExpectedNumber, token.Text)
				lx.finishStatement()
			}
			break
		case readParameterClose:
			switch token.Text {
			case ")":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.ExclusiveClose,
				}
				lx.state = readUnitEnd
				break
			case "]":
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.InclusiveClose,
				}
				lx.state = readUnitEnd
				break
			default:
				lx.err(token, errors.ExpectedRangeClose, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readUnaryStatementName:
			switch token.Text {
			case ";":
				lx.finishStatement()
			case ".":
				l := sc.lexemes[len(sc.lexemes)-1]
				if l.Type == lexer.Identifier {
					lexeme = &lexer.Lexeme{
						Type:    lexer.KeywordLiteral,
						Keyword: lexer.DereferenceKeyword,
					}
				} else {
					lx.err(token, errors.ExpectedIdentifier, ".")
					lx.finishStatement()
				}
				break
			default:
				l := sc.lexemes[len(sc.lexemes)-1]
				if l.Type != lexer.Identifier {
					if isIdentifier(token.Text) {
						lexeme = &lexer.Lexeme{
							Type: lexer.Identifier,
							Name: token.Text,
						}
					}
				} else {
					lx.err(token, errors.UnexpectedIdentifier, token.Text)
					lx.finishStatement()
				}
				break
			}
			break
		case readRequirementName:
			if isIdentifier(token.Text) {
				lexeme = &lexer.Lexeme{
					Type: lexer.Identifier,
					Name: token.Text,
				}
				lx.state = readRequirementCondition
			} else {
				lx.err(token, errors.ExpectedIdentifier, token.Text)
				lx.finishStatement()
			}
			break
		case readRequirementCondition:
			switch token.Text {
			case "!", "<", ">", "=":
				lx.tmp = token.Text[0]
				lx.state = readRequirementCondition2
				break
			default:
				lx.err(token, errors.ExpectedCondition, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readRequirementCondition2:
			switch lx.tmp {
			case '!':
				if token.Text == "=" {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.NotEqual,
					}
					lx.state = readExpression
				} else {
					lx.err(token, errors.ExpectedCondition, token.Text)
					lx.finishStatement()
				}
				break
			case '<':
				if token.Text == "=" {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.LessThanOrEqual,
					}
					lx.state = readExpression
				} else {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.LessThan,
					}
					lx.state = readExpression
					repeat = true
				}
				break
			case '>':
				if token.Text == "=" {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.GreaterThanOrEqual,
					}
					lx.state = readExpression
				} else {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.GreaterThan,
					}
					lx.state = readExpression
					repeat = true
				}
				break
			case '=':
				if token.Text == "=" {
					lexeme = &lexer.Lexeme{
						Type:      lexer.Conditional,
						Condition: lexer.GreaterThanOrEqual,
					}
					lx.state = readExpression
				} else {
					lx.err(token, errors.ExpectedCondition, token.Text)
					lx.finishStatement()
				}
				break
			default:
				lx.err(token, errors.MissingCase)
				break
			}
			break
		case readIfLHS:
			switch lx.expr.Read(token) {
			case slValid:
				break
			case slComplete:
				lx.state = readIfCondition
				repeat = true
				break
			case slError:
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.MissingCase)
				lx.finishStatement()
				break
			}
			break
		case readIfCondition:
			switch token.Text {
			case "!", "<", ">", "=":
				lx.tmp = token.Text[0]
				lx.state = readIfCondition2
				break
			default:
				lx.err(token, errors.ExpectedCondition, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readIfCondition2:
			switch lx.tmp {
			case '!':
				if token.Text == "=" {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.NotEqual
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
				} else {
					lx.err(token, errors.ExpectedCondition, token.Text)
					lx.finishStatement()
				}
				break
			case '<':
				if token.Text == "=" {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.LessThanOrEqual
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
				} else {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.LessThan
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
					repeat = true
				}
				break
			case '>':
				if token.Text == "=" {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.GreaterThanOrEqual
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
				} else {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.GreaterThan
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
					repeat = true
				}
				break
			case '=':
				if token.Text == "=" {
					l := sc.lexemes[len(sc.lexemes)-1]
					s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
					s.Operator = lexer.GreaterThanOrEqual
					lx.expr.Reset(&s.RHS)
					lx.state = readIfRHS
				} else {
					lx.err(token, errors.ExpectedCondition, token.Text)
					lx.finishStatement()
				}
				break
			default:
				lx.err(token, errors.MissingCase)
				break
			}
			break
		case readIfRHS:
			switch lx.expr.Read(token) {
			case slValid:
				break
			case slComplete:
				lx.state = readBlockStart
				repeat = true
				break
			case slError:
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.MissingCase)
				lx.finishStatement()
				break
			}
			break
		case readElse:
			switch token.Text {
			case "if":
				l := sc.lexemes[len(sc.lexemes)-1]
				l.SwitchBlocks = append(l.SwitchBlocks, lexer.SwitchBlock{
					LineNo:   token.LineNo,
					CharNo:   token.CharNo,
					FileName: token.FileName,
				})
				s := &l.SwitchBlocks[len(l.SwitchBlocks)-1]
				lx.expr.Reset(&s.LHS)
				lx.state = readIfLHS
				break
			case "{":
				lx.startScope()
				break
			default:
				lx.err(token, errors.ExpectedBlockStart, token.Text)
				lx.finishStatement()
				break
			}
			break
		case readPropertyLHS:
			l := sc.lexemes[len(sc.lexemes)-1]
			if l.Type != lexer.Expression {
				lexeme = &lexer.Lexeme{
					Type:       lexer.Expression,
					Expression: make([]lexer.ExpressionUnit, 0),
				}
				lx.expr.Reset(&lexeme.Expression)
			}
			switch lx.expr.Read(token) {
			case slValid:
				break
			case slComplete:
				if token.Text != "=" {
					lx.err(token, errors.ExpectedEquals, token.Text)
				}
				lexeme = &lexer.Lexeme{
					Type:    lexer.KeywordLiteral,
					Keyword: lexer.EqualsKeyword,
				}
				lx.state = readExpression
				break
			case slError:
				lx.finishStatement()
				break
			default:
				lx.err(token, errors.MissingCase)
				lx.finishStatement()
				break
			}
		default:
			lx.err(token, errors.StateError, lx.state, "lexerState")
			lx.state = statementStart
			break
		}
		if lexeme != nil {
			lexeme.LineNo = token.LineNo
			lexeme.CharNo = token.CharNo
			lexeme.FileName = token.FileName
			sc.lexemes = append(sc.lexemes, lexeme)
		}
	}
}

// EndStream is called by the tokenizer once one of the streams has finished
// being tokenized and streamed into the lexer.
func (lx *Lexer) EndStream(id int) {
	if id == 0 {
		lx.parser.End()
	}
}
