package auth

import (
	"fmt"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/cloud/rest"
)

// AuthorizationRulesGetter is the getter for authorization rules
type AuthorizationRulesGetter interface {
	AuthorizationRules(organizationID string) AuthorizationRuleInterface
}

// AuthorizationRuleInterface is the interface for authorization rules
type AuthorizationRuleInterface interface {
	Get(id string) (*types.AuthorizationRule, error)
	Delete(id string) error
	List() ([]types.AuthorizationRule, error)
	ListForRole(roleID string) ([]types.AuthorizationRule, error)
}

// authorizationRules implements AuthorizationRuleInterface
type authorizationRules struct {
	client         rest.Interface
	organizationID string
}

func newAuthorizationRules(c *Client, organizationID string) *authorizationRules {
	return &authorizationRules{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Get gets a authorization rule
func (c *authorizationRules) Get(id string) (*types.AuthorizationRule, error) {
	path := fmt.Sprintf("/v3/organizations/%s/rules/%s", c.organizationID, id)
	var out types.AuthorizationRule
	return &out, c.client.Get(path, &out)
}

// Delete deletes a authorization rule
func (c *authorizationRules) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/rules/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all authorization rules
func (c *authorizationRules) List() ([]types.AuthorizationRule, error) {
	path := fmt.Sprintf("/v3/organizations/%s/rules", c.organizationID)
	out := make([]types.AuthorizationRule, 0)
	return out, c.client.Get(path, &out)
}

// List lists all authorization rules for the given role
func (c *authorizationRules) ListForRole(roleID string) ([]types.AuthorizationRule, error) {
	path := fmt.Sprintf("/v3/organizations/%s/roles/%s/rules", c.organizationID, roleID)
	out := make([]types.AuthorizationRule, 0)
	return out, c.client.Get(path, &out)
}
