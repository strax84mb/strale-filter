package filter

type UnaryOperator struct {
	useOn   []OperatorPlacement
	operand Node
}

func (up *UnaryOperator) parse(p *Parser, index int, runes []rune, depth byte) (int, ErrorType) {
	var token string
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
			token, index = parseToken(index+1, runes, isOperatorRune)
			of, found := p.UnaryOperators[token]
			if !found {
				return index, ErrUnknownOperator
			}
			operator := &UnaryOperator{
				useOn: of.Placements,
			}
			up.operand = operator
			return operator.parse(p, index, runes, depth+1)
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
	return index + 1, NoError
}

func (up *UnaryOperator) GetType() NodeType {
	return UnaryOperatorType
}

func (up *UnaryOperator) Next() Node {
	return up.operand
}
