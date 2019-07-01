package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// UsersGetter is the getter for users
type UsersGetter interface {
	Users(organizationID string) UserInterface
}

// UserInterface is the interface for users
type UserInterface interface {
	Create(*types.User) (*types.User, error)
	Get(id string) (*types.User, error)
	Delete(id string) error
	List() ([]types.User, error)

	WithSSHAccess(clusterID string) UserInterface
}

// users implements UserInterface
type users struct {
	client         rest.Interface
	organizationID string
	sshClusterID   string
}

func newUsers(c *Client, organizationID string) *users {
	return &users{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a user
func (c *users) Create(*types.User) (*types.User, error) {
	// TODO
	return nil, nil
}

// Get gets a user
func (c *users) Get(id string) (*types.User, error) {
	path := fmt.Sprintf("/v3/organizations/%s/users/%s", c.organizationID, id)
	var out types.User
	return &out, c.client.Get(path, &out)
}

// Delete deletes a user
func (c *users) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/users/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all users. If the WithSSHAccess() function was called on this
// interface with a non-empty cluster ID, then only the users with SSH access
// to the given cluster will be listed.
func (c *users) List() ([]types.User, error) {
	var path string
	if c.sshClusterID != "" {
		path = fmt.Sprintf("/v3/organizations/%s/clusters/%s/ssh-users", c.organizationID, c.sshClusterID)
	} else {
		path = fmt.Sprintf("/v3/organizations/%s/users", c.organizationID)
	}
	out := make([]types.User, 0)
	return out, c.client.Get(path, &out)
}

// WithSSHAccess limits the users returned to only those that have SSH access
// to the given cluster. This function should only be used by a client that
// uses a cluster token.
func (c *users) WithSSHAccess(clusterID string) UserInterface {
	c.sshClusterID = clusterID
	return c
}
