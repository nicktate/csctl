package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/auth/types"
)

var (
	roleBindingTime = "1517001176920"

	roleBindings = []types.AuthorizationRoleBinding{
		{
			ID:                  types.UUID("1234"),
			Type:                strptr(types.AuthorizationRoleBindingTypeUserBinding),
			AuthorizationRoleID: types.UUID("1234"),
			UserID:              types.UUID("1234"),
			OwnerID:             types.UUID("1234"),
			CreatedAt:           &roleBindingTime,
		},
		{
			ID:                  types.UUID("4321"),
			Type:                strptr(types.AuthorizationRoleBindingTypeTeamBinding),
			AuthorizationRoleID: types.UUID("3456"),
			TeamID:              types.UUID("3456"),
			OwnerID:             types.UUID("4321"),
			CreatedAt:           &roleBindingTime,
		},
	}

	roleBindingSingle = []types.AuthorizationRoleBinding{roleBindings[0]}
)

func TestAuthorizationRoleBindingsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRoleBindings(roleBindingSingle)
	err := cluster.JSON(buf, true)
	assert.Nil(t, err)

	err = cluster.JSON(buf, false)
	assert.Nil(t, err)
}

func TestAuthorizationRoleBindingsYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRoleBindings(roleBindingSingle)
	err := cluster.YAML(buf, true)
	assert.Nil(t, err)

	err = cluster.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewAuthorizationRoleBindings(t *testing.T) {
	a := NewAuthorizationRoleBindings(nil)
	assert.NotNil(t, a)

	a = NewAuthorizationRoleBindings(roleBindings)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(roleBindings))

	a = AuthorizationRoleBinding()
	assert.NotNil(t, a)
}

func TestAuthorizationRoleBindingsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAuthorizationRoleBindings(roleBindings)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(roleBindings), info.numRows)
}
