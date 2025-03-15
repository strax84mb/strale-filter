package filter

type UnaryOperator struct {
	useOn   []OperatorPlacement
	operand Node
}

func (up *UnaryOperator) parse(p *Parser, index int, runes []rune, depth byte) (int, ErrorType) {
	switch runes[index] {
	case '(':
		if AllowPlacement(up.useOn, OnCondition) {
			// TODO handle condition
			panic(`"condition" not supported yet`)
		} else {
			return index, ErrUnexpectedToken
		}
	case '-':
		if AllowPlacement(up.useOn, OnUnaryOperator) {
			// TODO handle unary operator
			panic(`"unary operator" not supported yet`)
		} else {
			return index, ErrUnexpectedToken
		}
	default:
		var token string
		token, index = parseToken(index, runes, isFieldRune)
		up.operand = &Field{
			field: token,
		}
	}
	if index == len(runes) {
		return index, NoError
	}
	if index > len(runes) || runes[index] != ')' {
		return index, ErrUnexpectedToken
	}
	return index, NoError
}
