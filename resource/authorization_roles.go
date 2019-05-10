package resource

import (
	"io"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// AuthorizationRoles is a list of the associated cloud resource with additional functionality
type AuthorizationRoles struct {
	resource
	filterable
	items []types.AuthorizationRole
}

// NewAuthorizationRoles constructs a new AuthorizationRoles wrapping the given cloud type
func NewAuthorizationRoles(items []types.AuthorizationRole) *AuthorizationRoles {
	return &AuthorizationRoles{
		resource: resource{
			name:   "role",
			plural: "roles",
		},
		items: items,
	}
}

// AuthorizationRole constructs a new AuthorizationRoles with no underlying items, useful for
// interacting with the metadata itself.
func AuthorizationRole() *AuthorizationRoles {
	return NewAuthorizationRoles(nil)
}

func (c *AuthorizationRoles) columns() []string {
	return []string{
		"ID",
		"Name",
		"Description",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *AuthorizationRoles) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, role := range c.items {
		description := emptyColState
		if role.Description != "" {
			description = role.Description
		}

		table.Append([]string{
			string(role.ID),
			*role.Name,
			description,
			string(role.OwnerID),
			convert.UnixTimeMSToString(*role.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *AuthorizationRoles) JSON(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayJSON(w, c.items[0])
	}

	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *AuthorizationRoles) YAML(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayYAML(w, c.items[0])
	}

	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *AuthorizationRoles) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (c *AuthorizationRoles) FilterByOwnerID(id string) {
	filtered := make([]types.AuthorizationRole, 0)
	for _, resource := range c.items {
		if string(resource.OwnerID) == id {
			filtered = append(filtered, resource)
		}
	}

	c.items = filtered
}
