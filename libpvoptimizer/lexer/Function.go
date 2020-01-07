package lexer

// Function represents a mathematical function
type Function int

const (
	// Sine function
	Sine Function = 0
	// Cosine function
	Cosine Function = 1
	// Tangent function
	Tangent Function = 2
	// Cosecant function
	Cosecant Function = 3
	// Secant function
	Secant Function = 4
	// Cotangent function
	Cotangent Function = 5
	// ArcSine function
	ArcSine Function = 6
	// ArcCosine function
	ArcCosine Function = 7
	// ArcTangent function
	ArcTangent Function = 8
	// ArcCosecant function
	ArcCosecant Function = 9
	// ArcSecant function
	ArcSecant Function = 10
	// ArcCotangent function
	ArcCotangent Function = 11
	// HyperbolicSine function
	HyperbolicSine Function = 12
	// HyperbolicCosine function
	HyperbolicCosine Function = 13
	// HyperbolicTangent function
	HyperbolicTangent Function = 14
	// HyperbolicCosecant function
	HyperbolicCosecant Function = 15
	// HyperbolicSecant function
	HyperbolicSecant Function = 16
	// HyperbolicCotangent function
	HyperbolicCotangent Function = 17
	// HyperbolicArcSine function
	HyperbolicArcSine Function = 18
	// HyperbolicArcCosine function
	HyperbolicArcCosine Function = 19
	// HyperbolicArcTangent function
	HyperbolicArcTangent Function = 20
	// HyperbolicArcCosecant function
	HyperbolicArcCosecant Function = 21
	// HyperbolicArcSecant function
	HyperbolicArcSecant Function = 22
	// HyperbolicArcCotangent function
	HyperbolicArcCotangent Function = 23
	// Parenthesis is represented as an identity function
	Parenthesis Function = 24
	// Exponential function
	Exponential Function = 25
	// Logarithm function (natural base)
	Logarithm Function = 26
)
