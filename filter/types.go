package filter

type NodeType uint8

const (
	RootType NodeType = iota
	ConditionType
	NameType
	ValueType
	ArrayType
	UnaryOperatorType
	BinaryOperatorType
	ParametersType
	ParametrizedOperatorType
)

type Node interface {
	parse(p *Parser, index int, runes []rune, depth byte) (int, ErrorType)
	GetType() NodeType
	Next() Node
}

type OperatorPlacement uint8

const (
	OnField OperatorPlacement = iota
	OnCondition
	OnUnaryOperator
)

type OperatorsDefinition struct {
	Placements      []OperatorPlacement
	AllowedOperands []NodeType
}

func presentInODArray[T NodeType | OperatorPlacement](elements []T, value T) bool {
	for _, v := range elements {
		if v == value {
			return true
		}
	}
	return false
}

func (od *OperatorsDefinition) AllowPlacement(placement OperatorPlacement) bool {
	return presentInODArray(od.Placements, placement)
}

func (od *OperatorsDefinition) AllowOperand(nodeType NodeType) bool {
	return presentInODArray(od.AllowedOperands, nodeType)
}

type Parser struct {
	UnaryOperators        map[string]OperatorsDefinition
	BinaryOperators       map[string]OperatorsDefinition
	ParametrizedOperators map[string]OperatorsDefinition
}

func AllowPlacement(allowedPlacements []OperatorPlacement, placement OperatorPlacement) bool {
	return presentInODArray(allowedPlacements, placement)
}

func AllowOperand(allowedOperands []NodeType, operand NodeType) bool {
	return presentInODArray(allowedOperands, operand)
}
