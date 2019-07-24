package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	etcdVersion   = "3.2.24"
	dockerVersion = "18.6.1"

	nps = []types.NodePool{
		{
			Name:              strptr("test1"),
			ID:                types.UUID("1234"),
			Os:                strptr("ubuntu"),
			KubernetesMode:    strptr("master"),
			Count:             int32ptr(3),
			KubernetesVersion: strptr("1.12.1"),
			EtcdVersion:       &etcdVersion,
			DockerVersion:     &dockerVersion,
			Autoscaling: &types.NodePoolAutoscaling{
				Enabled: boolptr(false),
			},
		},
		{
			Name:              strptr("test2"),
			ID:                types.UUID("4321"),
			Os:                strptr("ubuntu"),
			KubernetesMode:    strptr("worker"),
			Count:             int32ptr(2),
			KubernetesVersion: strptr("1.11.1"),
			EtcdVersion:       nil,
			DockerVersion:     &dockerVersion,
			Autoscaling: &types.NodePoolAutoscaling{
				Enabled: boolptr(true),
			},
		},
	}

	npsSingle = []types.NodePool{
		{
			Name:              strptr("test3"),
			ID:                types.UUID("1234"),
			Os:                strptr("centos"),
			KubernetesMode:    strptr("master"),
			Count:             int32ptr(1),
			KubernetesVersion: strptr("1.12.1"),
			EtcdVersion:       &etcdVersion,
			DockerVersion:     &dockerVersion,
			Autoscaling: &types.NodePoolAutoscaling{
				Enabled: boolptr(false),
			},
		},
	}
)

func TestNodePoolsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodePools(npsSingle)
	err := a.JSON(buf, true)
	assert.Nil(t, err)

	err = a.JSON(buf, false)
	assert.Nil(t, err)
}

func TestNodePoolsYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodePools(npsSingle)
	err := a.YAML(buf, true)
	assert.Nil(t, err)

	err = a.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewNodePools(t *testing.T) {
	a := NewNodePools(nil)
	assert.NotNil(t, a)

	a = NewNodePools(nps)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(nps))

	a = NodePool()
	assert.NotNil(t, a)
}

func TestNodePoolsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewNodePools(nps)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(nps), info.numRows)
}
