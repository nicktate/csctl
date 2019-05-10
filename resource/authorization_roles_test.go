package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/auth/types"
)

var (
	roleTime = "1517001176920"

	roles = []types.AuthorizationRole{
		{
			ID:          types.UUID("1234"),
			Name:        strptr("role1"),
			Description: "desc1",
			OwnerID:     types.UUID("1234"),
			CreatedAt:   &roleTime,
		},
		{
			ID:          types.UUID("4321"),
			Name:        strptr("role2"),
			Description: "desc2",
			OwnerID:     types.UUID("4321"),
			CreatedAt:   &roleTime,
		},
	}

	roleSingle = []types.AuthorizationRole{roles[0]}
)

func TestAuthorizationRolesJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRoles(roleSingle)
	err := cluster.JSON(buf, true)
	assert.Nil(t, err)

	err = cluster.JSON(buf, false)
	assert.Nil(t, err)
}

func TestAuthorizationRolesYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRoles(roleSingle)
	err := cluster.YAML(buf, true)
	assert.Nil(t, err)

	err = cluster.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewAuthorizationRoles(t *testing.T) {
	a := NewAuthorizationRoles(nil)
	assert.NotNil(t, a)

	a = NewAuthorizationRoles(roles)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(roles))

	a = AuthorizationRole()
	assert.NotNil(t, a)
}

func TestAuthorizationRolesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAuthorizationRoles(roles)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(roles), info.numRows)
}
