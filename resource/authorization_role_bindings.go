package resource

import (
	"io"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// AuthorizationRoleBindings is a list of the associated cloud resource with additional functionality
type AuthorizationRoleBindings struct {
	resource
	filterable
	items []types.AuthorizationRoleBinding
}

// NewAuthorizationRoleBindings constructs a new AuthorizationRoleBindings wrapping the given cloud type
func NewAuthorizationRoleBindings(items []types.AuthorizationRoleBinding) *AuthorizationRoleBindings {
	return &AuthorizationRoleBindings{
		resource: resource{
			name:   "role-binding",
			plural: "role-bindings",
		},
		items: items,
	}
}

// AuthorizationRoleBinding constructs a new AuthorizationRoleBindings with no underlying items, useful for
// interacting with the metadata itself.
func AuthorizationRoleBinding() *AuthorizationRoleBindings {
	return NewAuthorizationRoleBindings(nil)
}

func (c *AuthorizationRoleBindings) columns() []string {
	return []string{
		"ID",
		"Type",
		"Role ID",
		"User ID",
		"Team ID",
		"Cluster ID",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *AuthorizationRoleBindings) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, binding := range c.items {
		userID := emptyColState
		if binding.UserID != "" {
			userID = string(binding.UserID)
		}
		teamID := emptyColState
		if binding.TeamID != "" {
			teamID = string(binding.TeamID)
		}
		clusterID := emptyColState
		if binding.ClusterID != "" {
			clusterID = string(binding.ClusterID)
		}

		table.Append([]string{
			string(binding.ID),
			*binding.Type,
			string(binding.AuthorizationRoleID),
			userID,
			teamID,
			clusterID,
			string(binding.OwnerID),
			convert.UnixTimeMSToString(*binding.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *AuthorizationRoleBindings) JSON(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayJSON(w, c.items[0])
	}

	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *AuthorizationRoleBindings) YAML(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayYAML(w, c.items[0])
	}

	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *AuthorizationRoleBindings) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (c *AuthorizationRoleBindings) FilterByOwnerID(id string) {
	filtered := make([]types.AuthorizationRoleBinding, 0)
	for _, resource := range c.items {
		if string(resource.OwnerID) == id {
			filtered = append(filtered, resource)
		}
	}

	c.items = filtered
}
