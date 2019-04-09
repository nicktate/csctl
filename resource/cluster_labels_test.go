package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	clusterLabelTime = "1517001176920"

	clusterLabels = []types.ClusterLabel{
		{
			ID:        types.UUID("1234"),
			Key:       strptr("key1"),
			Value:     strptr("value1"),
			CreatedAt: clusterLabelTime,
			UpdatedAt: clusterLabelTime,
		},
		{
			ID:        types.UUID("4321"),
			Key:       strptr("key2"),
			Value:     strptr("value2"),
			CreatedAt: clusterLabelTime,
			UpdatedAt: clusterLabelTime,
		},
	}

	clusterLabelsSingle = []types.ClusterLabel{
		{
			ID:        types.UUID("1234"),
			Key:       strptr("key3"),
			Value:     strptr("value3"),
			CreatedAt: clusterLabelTime,
			UpdatedAt: clusterLabelTime,
		},
	}
)

func TestNewClusterLabels(t *testing.T) {
	a := NewNodes(nil)
	assert.NotNil(t, a)

	a = NewNodes(nodes)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(nodes))

	a = Node()
	assert.NotNil(t, a)
}

func TestClusterLabelsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewClusterLabels(clusterLabels)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(nodes), info.numRows)
}

func TestBuildClusterLabelString(t *testing.T) {
	type stringTest struct {
		label    types.ClusterLabel
		expected string
	}

	var tests = []stringTest{
		{
			label: types.ClusterLabel{
				Key:   strptr("nodepool.containership.io/key1"),
				Value: strptr("value1"),
			},
			expected: "nodepool.containership.io/key1=value1",
		},
		{
			label: types.ClusterLabel{
				Key:   strptr("nodepool.containership.io/nil-value"),
				Value: nil,
			},
			expected: "nodepool.containership.io/nil-value=",
		},
		{
			label: types.ClusterLabel{
				Key:   strptr("nodepool.containership.io/empty-value"),
				Value: nil,
			},
			expected: "nodepool.containership.io/empty-value=",
		},
	}

	for _, test := range tests {
		result := buildClusterLabelString(test.label)
		assert.Equal(t, test.expected, result)
	}
}
