package filter

func (p *Parser) Parse(text string) (Node, error) {
	runes := []rune(text)
	rn := &rootNode{}
	err := rn.parse(p, 0, runes, 0)
	return rn, err
}
