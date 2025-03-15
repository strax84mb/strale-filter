package filter

type Field struct {
	field string
}

func (f *Field) parse(p *Parser, index int, runes []rune, depth byte) (int, ErrorType) {
	return index, NoError
}

func (f *Field) GetType() NodeType {
	return NameType
}

func (f *Field) Next() Node {
	return nil
}
