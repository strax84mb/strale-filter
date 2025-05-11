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

func TestNotExists(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse("-not(-exists(table.name))")
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestFieldNumeral(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse("table.num -eq 12")
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestFieldTrue(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse("table.bool -eq true")
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestFieldFalse(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse("table.bool -eq false")
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestNotNakedCondition(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse(`-not(table.name -eq "qwerty")`)
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestNotCondition(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse(`-not((table.name -eq "qwerty"))`)
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestNotConditions(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse(`-not((table.name -eq "qwerty") -and (table.type -not(-eq("qwerty"))))`)
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestInConditions(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse(`table.name -in ["val_1", "val_2", "val_3"]`)
	assert.NoError(t, err)
	assert.NotNil(t, node)
}

func TestMultiConditions(t *testing.T) {
	parser := makeParser()
	node, err := parser.Parse(`(table.name -eq "name") -and (table.type -eq "some type") -and (table.date -eq "2025-03-12") -or (table.num -eq 12)`)
	assert.NoError(t, err)
	assert.NotNil(t, node)
}
