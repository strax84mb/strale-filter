package filter

import (
	"fmt"
)

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

func (n *rootNode) parse(p *Parser, index int, runes []rune, depth byte) (int, error) {
	r := runes[index]
	switch {
	case r == '(':
		// recognize condition
		panic(`"(" is not supported yet`)
	case r == '-':
		return n.parseUnaryOperator(p, index+1, runes, depth)
	case isFieldRune(r):
		// recognize field
	default:
	}
	return 0, nil
}

func (n *rootNode) parseUnaryOperator(p *Parser, index int, runes []rune, depth byte) (int, error) {
	operatorName, nextCharIndex := parseToken(index, runes, isOperatorRune)
	of, found := p.UnaryOperators[operatorName]
	if !found {
		return index, fmt.Errorf("first character error: %w", ErrUnknownOperator)
	}
	operator := UnaryOperator{
		useOn: of.Placements,
	}
	var err error
	nextCharIndex, err = operator.parse(p, nextCharIndex, runes, depth+1)
	if err != nil {
		return nextCharIndex, err
	}
	if nextCharIndex != len(runes) {
		return nextCharIndex, ErrNotExhausted
	}
	return nextCharIndex, nil
}
