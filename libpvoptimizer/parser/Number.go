package parser

import "strconv"

// Number is a numerical quantity
type Number float64

// ParseNumber parses a string into a Number
func ParseNumber(str string) (Number, error) {
	f, e := strconv.ParseFloat(str, 64)
	return Number(f), e
}
