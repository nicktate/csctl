package resource

import (
	"io"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// AuthorizationRules is a list of the associated cloud resource with additional functionality
type AuthorizationRules struct {
	resource
	filterable
	items []types.AuthorizationRule
}

// NewAuthorizationRules constructs a new AuthorizationRules wrapping the given cloud type
func NewAuthorizationRules(items []types.AuthorizationRule) *AuthorizationRules {
	return &AuthorizationRules{
		resource: resource{
			name:   "rule",
			plural: "rules",
		},
		items: items,
	}
}

// AuthorizationRule constructs a new AuthorizationRules with no underlying items, useful for
// interacting with the metadata itself.
func AuthorizationRule() *AuthorizationRules {
	return NewAuthorizationRules(nil)
}

func (c *AuthorizationRules) columns() []string {
	return []string{
		"ID",
		"Name",
		"Description",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *AuthorizationRules) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, rule := range c.items {
		description := emptyColState
		if rule.Description != "" {
			description = rule.Description
		}

		table.Append([]string{
			string(rule.ID),
			*rule.Name,
			description,
			string(rule.OwnerID),
			convert.UnixTimeMSToString(*rule.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *AuthorizationRules) JSON(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayJSON(w, c.items[0])
	}

	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *AuthorizationRules) YAML(w io.Writer, listView bool) error {
	if !listView && len(c.items) == 1 {
		return displayYAML(w, c.items[0])
	}

	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *AuthorizationRules) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (c *AuthorizationRules) FilterByOwnerID(id string) {
	filtered := make([]types.AuthorizationRule, 0)
	for _, resource := range c.items {
		if string(resource.OwnerID) == id {
			filtered = append(filtered, resource)
		}
	}

	c.items = filtered
}
