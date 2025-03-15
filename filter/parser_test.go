package filter_test

import (
	"testing"

	"github.com/strax84mb/strale-filter/filter"
	"github.com/stretchr/testify/assert"
)

func makeParser() *filter.Parser {
	return &filter.Parser{
		UnaryOperators: map[string]filter.OperatorsDefinition{
			"not": {
				Placements: []filter.OperatorPlacement{filter.OnCondition, filter.OnUnaryOperator},
			},
			"exists": {
				Placements: []filter.OperatorPlacement{filter.OnField},
			},
		},
		BinaryOperators: map[string]filter.OperatorsDefinition{
			"eq": {
				Placements:      []filter.OperatorPlacement{filter.OnField},
				AllowedOperands: []filter.NodeType{filter.ValueType},
			},
			"in": {
				Placements:      []filter.OperatorPlacement{filter.OnField},
				AllowedOperands: []filter.NodeType{filter.ArrayType},
			},
			"and": {
				Placements:      []filter.OperatorPlacement{filter.OnCondition, filter.OnUnaryOperator},
				AllowedOperands: []filter.NodeType{filter.ConditionType, filter.UnaryOperatorType},
			},
			"or": {
				Placements:      []filter.OperatorPlacement{filter.OnCondition, filter.OnUnaryOperator},
				AllowedOperands: []filter.NodeType{filter.ConditionType, filter.UnaryOperatorType},
			},
		},
	}
}

func TestExists(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse("-exists(table.name)")
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestNotExists(t *testing.T) {}

func TestNotCondition(t *testing.T) {}
