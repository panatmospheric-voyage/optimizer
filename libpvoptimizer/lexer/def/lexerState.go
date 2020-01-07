package def

type lexerState int

const (
	statementStart            lexerState = 0
	readUnitLHS               lexerState = 1
	readUnitEquals            lexerState = 2
	readUnitRHS               lexerState = 3
	readUnitEnd               lexerState = 4
	readEndOfStatement        lexerState = 5
	readPropertyName          lexerState = 6
	readPropertyEquals        lexerState = 7
	readExpression            lexerState = 8
	readBlockName             lexerState = 9
	readBlockStart            lexerState = 10
	readParameterName         lexerState = 11
	readParameterOpen         lexerState = 12
	readParameterLower        lexerState = 13
	readParameterDelim        lexerState = 14
	readParameterUpper        lexerState = 15
	readParameterClose        lexerState = 16
	readUnaryStatementName    lexerState = 17
	readRequirementName       lexerState = 18
	readRequirementCondition  lexerState = 19
	readRequirementCondition2 lexerState = 20
	readIfLHS                 lexerState = 21
	readIfCondition           lexerState = 22
	readIfCondition2          lexerState = 23
	readIfRHS                 lexerState = 24
	readElse                  lexerState = 25
	readPropertyLHS           lexerState = 26
)
