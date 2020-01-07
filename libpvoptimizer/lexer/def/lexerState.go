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
	readRequirementName       lexerState = 17
	readRequirementCondition  lexerState = 18
	readRequirementCondition2 lexerState = 19
	readIfLHS                 lexerState = 20
	readIfCondition           lexerState = 21
	readIfCondition2          lexerState = 22
	readIfRHS                 lexerState = 23
	readElse                  lexerState = 24
	readPropertyLHS           lexerState = 25
)
