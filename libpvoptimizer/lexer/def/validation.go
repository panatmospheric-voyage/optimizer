package def

func isIdentifier(str string) bool {
	f := str[0]
	return (f >= 'a' && f <= 'z') || (f >= 'A' && f <= 'Z') || f == '_'
}

func isNumber(str string) bool {
	f := str[0]
	return (f >= '0' && f <= '9') || (f == '-' && len(str) > 1) || f == '.'
}
