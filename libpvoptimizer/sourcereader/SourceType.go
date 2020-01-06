package sourcereader

// SourceType represents the location of an included source file.
type SourceType int

const (
	// DefaultSource means it is the main file that is being optimized
	DefaultSource SourceType = 0
	// UserSource means it is an included file that the user wrote
	UserSource SourceType = 1
	// SystemSource means it is an included file in the system standard library
	SystemSource SourceType = 2
)
