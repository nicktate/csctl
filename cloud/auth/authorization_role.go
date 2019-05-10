package auth

import (
	"fmt"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/cloud/rest"
)

// AuthorizationRolesGetter is the getter for authorization roles
type AuthorizationRolesGetter interface {
	AuthorizationRoles(organizationID string) AuthorizationRoleInterface
}

// AuthorizationRoleInterface is the interface for authorization roles
type AuthorizationRoleInterface interface {
	Get(id string) (*types.AuthorizationRole, error)
	Delete(id string) error
	List() ([]types.AuthorizationRole, error)
}

// authorizationRoles implements AuthorizationRoleInterface
type authorizationRoles struct {
	client         rest.Interface
	organizationID string
}

func newAuthorizationRoles(c *Client, organizationID string) *authorizationRoles {
	return &authorizationRoles{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Get gets a authorization role
func (c *authorizationRoles) Get(id string) (*types.AuthorizationRole, error) {
	path := fmt.Sprintf("/v3/organizations/%s/roles/%s", c.organizationID, id)
	var out types.AuthorizationRole
	return &out, c.client.Get(path, &out)
}

// Delete deletes a authorization role
func (c *authorizationRoles) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/roles/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all authorization roles
func (c *authorizationRoles) List() ([]types.AuthorizationRole, error) {
	path := fmt.Sprintf("/v3/organizations/%s/roles", c.organizationID)
	out := make([]types.AuthorizationRole, 0)
	return out, c.client.Get(path, &out)
}
