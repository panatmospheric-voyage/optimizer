package def

type expressionState int

const (
	readVariable        expressionState = 0
	readOperator        expressionState = 1
	readOpenParenthesis expressionState = 2
	readVariableDot     expressionState = 3
	readVariableName    expressionState = 4
	readOperatorOrUnit  expressionState = 5
	readUnit            expressionState = 6
)
