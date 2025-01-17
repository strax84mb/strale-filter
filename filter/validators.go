package filter

func isLowerCase(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isFieldRune(r rune) bool {
	return isLowerCase(r) || isUpperCase(r) || isDigit(r) || r == '.' || r == '-' || r == '_'
}

func isOperatorRune(r rune) bool {
	return isLowerCase(r) || isUpperCase(r) || isDigit(r)
}
