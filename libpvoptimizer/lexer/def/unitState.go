package def

type unitState int

const (
	readUnitName   unitState = 0
	afterUnitName  unitState = 1
	readUnitPower  unitState = 2
	afterUnitPower unitState = 3
)
