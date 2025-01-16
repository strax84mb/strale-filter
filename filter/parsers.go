package filter

func Parse(text string) (Node, error) {
	runes := []rune(text)
	rn := &rootNode{}
	err := rn.parse(nil, 0, runes)
	return rn, err
}
