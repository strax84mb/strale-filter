package filter

var _ Node = &rootNode{}

type rootNode struct {
	next Node
}

// Next implements Node.
func (n *rootNode) Next() Node {
	return n.next
}

// GetType implements Node.
func (n *rootNode) GetType() NodeType {
	return RootType
}

func (n *rootNode) parse(p *Parser, index int, runes []rune, depth byte) (int, ErrorType) {
	var (
		err ErrorType
		r   rune = runes[index]
	)
	switch {
	case r == '(':
		// TODO recognize condition
		panic(`"(" is not supported yet`)
	case r == '-':
		index, err = n.parseUnaryOperator(p, index+1, runes, depth)
	case isFieldRune(r):
		// TODO recognize field
		panic(`"field" is not supported yet`)
	default:
	}
	if err != NoError {
		return index, err
	} else if index < len(runes) {
		return index, ErrNotExhausted
	}
	return 0, NoError
}

func (n *rootNode) parseUnaryOperator(p *Parser, index int, runes []rune, depth byte) (int, ErrorType) {
	operatorName, nextCharIndex := parseToken(index, runes, isOperatorRune)
	of, found := p.UnaryOperators[operatorName]
	if !found {
		return index, ErrUnknownOperator
	}
	operator := UnaryOperator{
		useOn: of.Placements,
	}
	var err ErrorType
	nextCharIndex, err = operator.parse(p, nextCharIndex, runes, depth+1)
	if err != NoError {
		return nextCharIndex, err
	}
	return nextCharIndex, NoError
}
