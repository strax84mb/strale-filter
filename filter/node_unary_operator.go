package filter

type UnaryOperator struct {
	useOn []OperatorPlacement
}

func (up *UnaryOperator) parse(p *Parser, index int, runes []rune, depth byte) (int, error) {
	return 0, nil
}
