package def

type sublexResult int

const (
	slValid    sublexResult = 0
	slComplete sublexResult = 1
	slError    sublexResult = 2
)
