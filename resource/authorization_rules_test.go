package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/auth/types"
)

var (
	ruleTime = "1517001176920"

	rules = []types.AuthorizationRule{
		{
			ID:          types.UUID("1234"),
			Name:        strptr("rule1"),
			Description: "desc1",
			OwnerID:     types.UUID("1234"),
			CreatedAt:   &ruleTime,
		},
		{
			ID:          types.UUID("4321"),
			Name:        strptr("rule2"),
			Description: "desc2",
			OwnerID:     types.UUID("4321"),
			CreatedAt:   &ruleTime,
		},
	}

	ruleSingle = []types.AuthorizationRule{rules[0]}
)

func TestAuthorizationRulesJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRules(ruleSingle)
	err := cluster.JSON(buf, true)
	assert.Nil(t, err)

	err = cluster.JSON(buf, false)
	assert.Nil(t, err)
}

func TestAuthorizationRulesYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	cluster := NewAuthorizationRules(ruleSingle)
	err := cluster.YAML(buf, true)
	assert.Nil(t, err)

	err = cluster.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewAuthorizationRules(t *testing.T) {
	a := NewAuthorizationRules(nil)
	assert.NotNil(t, a)

	a = NewAuthorizationRules(rules)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(rules))

	a = AuthorizationRule()
	assert.NotNil(t, a)
}

func TestAuthorizationRulesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAuthorizationRules(rules)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(rules), info.numRows)
}
