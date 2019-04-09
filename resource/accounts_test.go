package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	acctTime = "1517001176920"

	accts = []types.Account{
		{
			Name:      strptr("test1"),
			ID:        types.UUID("1234"),
			CreatedAt: &acctTime,
		},
		{
			Name:      strptr("test2"),
			ID:        types.UUID("4321"),
			CreatedAt: &acctTime,
		},
	}

	acctSingle = []types.Account{
		{
			Name:      strptr("test3"),
			ID:        types.UUID("1234"),
			CreatedAt: &acctTime,
		},
	}
)

func TestAccountsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewAccounts(acctSingle)
	err := a.JSON(buf, true)
	assert.Nil(t, err)

	err = a.JSON(buf, false)
	assert.Nil(t, err)
}

func TestAccountsYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewAccounts(acctSingle)
	err := a.YAML(buf, true)
	assert.Nil(t, err)

	err = a.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewAccounts(t *testing.T) {
	a := NewAccounts(nil)
	assert.NotNil(t, a)

	a = NewAccounts(accts)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(accts))

	a = Account()
	assert.NotNil(t, a)
}

func TestAccountsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAccounts(accts)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(accts), info.numRows)
}
