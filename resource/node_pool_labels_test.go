package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	nodePoolLabelTime = "1517001176920"

	nodePoolLabels = []types.NodePoolLabel{
		{
			ID:        types.UUID("1234"),
			Key:       strptr("key1"),
			Value:     strptr("value1"),
			CreatedAt: nodePoolLabelTime,
			UpdatedAt: nodePoolLabelTime,
		},
		{
			ID:        types.UUID("4321"),
			Key:       strptr("key2"),
			Value:     strptr("value2"),
			CreatedAt: nodePoolLabelTime,
			UpdatedAt: nodePoolLabelTime,
		},
	}

	nodePoolLabelsSingle = []types.NodePoolLabel{
		{
			ID:        types.UUID("1234"),
			Key:       strptr("key3"),
			Value:     strptr("value3"),
			CreatedAt: nodePoolLabelTime,
			UpdatedAt: nodePoolLabelTime,
		},
	}
)

func TestNewNodePoolLabels(t *testing.T) {
	a := NewNodes(nil)
	assert.NotNil(t, a)

	a = NewNodes(nodes)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(nodes))

	a = Node()
	assert.NotNil(t, a)
}

func TestNodePoolLabelsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodePoolLabels(nodePoolLabelsSingle)
	err := a.JSON(buf, true)
	assert.Nil(t, err)

	err = a.JSON(buf, false)
	assert.Nil(t, err)
}

func TestNodePoolLabelsYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodePoolLabels(nodePoolLabelsSingle)
	err := a.YAML(buf, true)
	assert.Nil(t, err)

	err = a.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNodePoolLabelsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewNodePoolLabels(nodePoolLabels)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(nodes), info.numRows)
}

func TestBuildNodePoolLabelString(t *testing.T) {
	type stringTest struct {
		label    types.NodePoolLabel
		expected string
	}

	var tests = []stringTest{
		{
			label: types.NodePoolLabel{
				Key:   strptr("nodepool.containership.io/key1"),
				Value: strptr("value1"),
			},
			expected: "nodepool.containership.io/key1=value1",
		},
		{
			label: types.NodePoolLabel{
				Key:   strptr("nodepool.containership.io/nil-value"),
				Value: nil,
			},
			expected: "nodepool.containership.io/nil-value=",
		},
		{
			label: types.NodePoolLabel{
				Key:   strptr("nodepool.containership.io/empty-value"),
				Value: nil,
			},
			expected: "nodepool.containership.io/empty-value=",
		},
	}

	for _, test := range tests {
		result := buildNodePoolLabelString(test.label)
		assert.Equal(t, test.expected, result)
	}
}
