package parser

import (
	"math"
	"strconv"
)

// Number is a numerical quantity
type Number float64

// ParseNumber parses a string into a Number
func ParseNumber(str string) (Number, error) {
	f, e := strconv.ParseFloat(str, 64)
	return Number(f), e
}

// Pow raises the number to a power
func (n Number) Pow(ex Number) Number {
	return Number(math.Pow(float64(n), float64(ex)))
}

// Sine function
func Sine(x Number) Number {
	return Number(math.Sin(float64(x)))
}

// Cosine function
func Cosine(x Number) Number {
	return Number(math.Cos(float64(x)))
}

// Tangent function
func Tangent(x Number) Number {
	return Number(math.Tan(float64(x)))
}

// ArcSine function
func ArcSine(x Number) Number {
	return Number(math.Asin(float64(x)))
}

// ArcCosine function
func ArcCosine(x Number) Number {
	return Number(math.Acos(float64(x)))
}

// ArcTangent function
func ArcTangent(x Number) Number {
	return Number(math.Atan(float64(x)))
}

// HyperbolicSine function
func HyperbolicSine(x Number) Number {
	return Number(math.Sinh(float64(x)))
}

// HyperbolicCosine function
func HyperbolicCosine(x Number) Number {
	return Number(math.Cosh(float64(x)))
}

// HyperbolicTangent function
func HyperbolicTangent(x Number) Number {
	return Number(math.Tanh(float64(x)))
}

// HyperbolicArcSine function
func HyperbolicArcSine(x Number) Number {
	return Number(math.Asinh(float64(x)))
}

// HyperbolicArcCosine function
func HyperbolicArcCosine(x Number) Number {
	return Number(math.Acosh(float64(x)))
}

// HyperbolicArcTangent function
func HyperbolicArcTangent(x Number) Number {
	return Number(math.Atanh(float64(x)))
}

// Exponential function
func Exponential(x Number) Number {
	return Number(math.Exp(float64(x)))
}

// Logarithm function
func Logarithm(x Number) Number {
	return Number(math.Log(float64(x)))
}

// Abs olute value function
func Abs(x Number) Number {
	return Number(math.Abs(float64(x)))
}
