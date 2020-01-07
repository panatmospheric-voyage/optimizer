package mock

import (
	"fmt"
	"strings"
	"sync"

	libpvoptimizer "../.."
	"../../errors"
	"../../lexer"
)

// Parser is the default implementation of IParser
type Parser struct {
	handler errors.IErrorHandler
}

// Init initializes the layer and is called from the pipeline layer
func (ps Parser) Init(lexer libpvoptimizer.ILexer, evaluator libpvoptimizer.IEvaluator, e errors.IErrorHandler, wg *sync.WaitGroup) {
	ps.handler = e
}

var (
	keywords = []string{
		"K_UNIT",
		"K_EQUALS",
		"K_PROPERTY",
		"K_ASSEMBLY",
		"K_PARAMETER",
		"K_SUMMARIZE",
		"K_ENUM",
		"K_VALUE",
		"K_INCLUSIVE_OPEN",
		"K_EXCLUSIVE_OPEN",
		"K_INCLUSIVE_CLOSE",
		"K_EXCLUSIVE_CLOSE",
		"K_ASSIGN_ARROW",
		"K_REQUIRE",
		"K_MINIMIZE",
		"K_MAXIMIZE",
		"K_COMMA",
	}
	operators = []string{
		"OP_ADD",
		"OP_SUB",
		"OP_MUL",
		"OP_DIV",
		"OP_EXP",
	}
	funcs = []string{
		"F_SIN",
		"F_COS",
		"F_TAN",
		"F_CSC",
		"F_SEC",
		"F_COT",
		"F_ASIN",
		"F_ACOS",
		"F_ATAN",
		"F_ACSC",
		"F_ASEC",
		"F_ACOT",
		"F_SINH",
		"F_COSH",
		"F_TANH",
		"F_CSCH",
		"F_SECH",
		"F_COTh",
		"F_ASINH",
		"F_ACOSH",
		"F_ATANH",
		"F_ACSCH",
		"F_ASECH",
		"F_ACOTH",
		"F_IDET",
		"F_ABS",
	}
	conditionals = []string{
		"C_LT",
		"C_LE",
		"C_GT",
		"C_GE",
		"C_EQ",
		"C_NE",
	}
)

func printExpression(expression []lexer.ExpressionUnit) {
	for i, expr := range expression {
		if i != 0 {
			fmt.Print(", ")
		}
		switch expr.Type {
		case lexer.ExpressionNumber:
			if len(expr.Unit) > 0 {
				fmt.Printf("#(%s ", expr.Text[0])
				for j, unit := range expr.Unit {
					if j != 0 {
						fmt.Print("*")
					}
					fmt.Printf("'%s'^%d", unit.Name, unit.Power)
				}
				fmt.Print(")")
			} else {
				fmt.Printf("#(%s)", expr.Text[0])
			}
			break
		case lexer.Variable:
			fmt.Printf("$(%s)", strings.Join(expr.Text, "."))
			break
		case lexer.OperatorSymbol:
			if expr.Operator >= 0 && int(expr.Operator) < len(operators) {
				fmt.Print(operators[expr.Operator])
			} else {
				fmt.Print("OP_?")
			}
			break
		case lexer.FunctionLiteral:
			if expr.Function >= 0 && int(expr.Function) < len(funcs) {
				fmt.Print(funcs[expr.Function])
			} else {
				fmt.Print("F_?")
			}
			fmt.Print("(")
			printExpression(expr.SubExpression)
			fmt.Print(")")
			break
		default:
			fmt.Printf("?")
			break
		}
	}
}

func printStatement(statement lexer.Statement) {
	for _, lexeme := range statement.Lexemes {
		fmt.Print("[")
		switch lexeme.Type {
		case lexer.KeywordLiteral:
			if lexeme.Keyword >= 0 && int(lexeme.Keyword) < len(keywords) {
				fmt.Print(keywords[lexeme.Keyword])
			} else {
				fmt.Print("K_?")
			}
			break
		case lexer.UnitLiteral:
			for i, unit := range lexeme.Unit {
				if i != 0 {
					fmt.Print(", ")
				}
				fmt.Printf("'%s'^%d", unit.Name, unit.Power)
			}
			break
		case lexer.NumberLiteral:
			fmt.Printf("#:%s", lexeme.Name)
			break
		case lexer.Expression:
			fmt.Print("e{")
			printExpression(lexeme.Expression)
			fmt.Print("}")
			break
		case lexer.GroupBlock:
			fmt.Printf("<%s>{", lexeme.Name)
			for i, st := range lexeme.Statements {
				if i != 0 {
					fmt.Print(", ")
				}
				printStatement(st)
			}
			fmt.Print("}")
			break
		case lexer.Identifier:
			fmt.Printf("<%s>", lexeme.Name)
			break
		case lexer.Switch:
			for i, sw := range lexeme.SwitchBlocks {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Print("if e{")
				printExpression(sw.LHS)
				fmt.Print("} ")
				if sw.Operator >= 0 && int(sw.Operator) < len(conditionals) {
					fmt.Print(conditionals[sw.Operator])
				} else {
					fmt.Print("C_?")
				}
				fmt.Print(" e{")
				printExpression(sw.RHS)
				fmt.Print("} then { ")
				for j, st := range sw.Statements {
					if j != 0 {
						fmt.Print(", ")
					}
					fmt.Print("<")
					printStatement(st)
					fmt.Print(">")
				}
				fmt.Print(" }")
			}
			break
		case lexer.Conditional:
			if lexeme.Condition >= 0 && int(lexeme.Condition) < len(conditionals) {
				fmt.Print(conditionals[lexeme.Condition])
			} else {
				fmt.Print("C_?")
			}
			break
		default:
			fmt.Print("?")
			break
		}
		fmt.Print("]")
	}
}

// Stream accepts a statement and processes it.
func (ps Parser) Stream(statement lexer.Statement) {
	printStatement(statement)
	fmt.Println()
}

// End is called by the lexer once all statements have been streamed into
// the parser.  This is when the model is sent into the evaluator layer.
func (ps Parser) End() {
	fmt.Println("End")
}
