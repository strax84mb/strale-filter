package filter

func (p *Parser) Parse(text string) (Node, error) {
	runes := []rune(text)
	rn := &rootNode{}
	i, err := rn.parse(p, 0, runes, 0)
	if err != NoError {
		return nil, newParseError(err, runes, i)
	}
	return rn, nil
}

func parseToken(index int, runes []rune, isRuneAllowed func(rune) bool) (string, int) {
	i := index
	end := len(runes)
	if i == end {
		return "", end
	}
	for {
		i++
		if i == end {
			return string(runes[index:]), end
		}
		if !isRuneAllowed(runes[i]) {
			break
		}
	}
	return string(runes[index:i]), i + 1
}
