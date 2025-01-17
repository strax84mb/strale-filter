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

func (n *rootNode) parse(p *Parser, ndex int, runes []rune, depth byte) (int, error) {
	return 0, nil
}
