package auth

import (
	"fmt"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/cloud/rest"
)

// AuthorizationRoleBindingsGetter is the getter for authorization role bindings
type AuthorizationRoleBindingsGetter interface {
	AuthorizationRoleBindings(organizationID string) AuthorizationRoleBindingInterface
}

// AuthorizationRoleBindingInterface is the interface for authorization role bindings
type AuthorizationRoleBindingInterface interface {
	Get(id string) (*types.AuthorizationRoleBinding, error)
	Delete(id string) error
	List() ([]types.AuthorizationRoleBinding, error)
	ListForRole(roleID string) ([]types.AuthorizationRoleBinding, error)
}

// authorizationRoleBindings implements AuthorizationRoleBindingInterface
type authorizationRoleBindings struct {
	client         rest.Interface
	organizationID string
}

func newAuthorizationRoleBindings(c *Client, organizationID string) *authorizationRoleBindings {
	return &authorizationRoleBindings{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Get gets a authorization role binding
func (c *authorizationRoleBindings) Get(id string) (*types.AuthorizationRoleBinding, error) {
	path := fmt.Sprintf("/v3/organizations/%s/role-bindings/%s", c.organizationID, id)
	var out types.AuthorizationRoleBinding
	return &out, c.client.Get(path, &out)
}

// Delete deletes a authorization role binding
func (c *authorizationRoleBindings) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/role-bindings/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all authorization role bindings
func (c *authorizationRoleBindings) List() ([]types.AuthorizationRoleBinding, error) {
	path := fmt.Sprintf("/v3/organizations/%s/role-bindings", c.organizationID)
	out := make([]types.AuthorizationRoleBinding, 0)
	return out, c.client.Get(path, &out)
}

// ListForRole lists all authorization role bindings for a specific role
func (c *authorizationRoleBindings) ListForRole(roleID string) ([]types.AuthorizationRoleBinding, error) {
	path := fmt.Sprintf("/v3/organizations/%s/roles/%s/role-bindings", c.organizationID, roleID)
	out := make([]types.AuthorizationRoleBinding, 0)
	return out, c.client.Get(path, &out)
}
