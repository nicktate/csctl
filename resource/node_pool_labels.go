package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/label"
	"github.com/containership/csctl/resource/table"
)

// NodePoolLabels is a list of the associated cloud resource with additional functionality
type NodePoolLabels struct {
	resource
	items []types.NodePoolLabel
}

// NewNodePoolLabels constructs a new NodePoolLabels wrapping the given cloud type
func NewNodePoolLabels(items []types.NodePoolLabel) *NodePoolLabels {
	return &NodePoolLabels{
		resource: resource{
			name:    "node-pool-label",
			plural:  "node-pool-labels",
			aliases: []string{"npl", "npls"},
		},
		items: items,
	}
}

// NodePoolLabel constructs a new NodePoolLabels with no underlying items, useful for
// interacting with the metadata itself.
func NodePoolLabel() *NodePoolLabels {
	return NewNodePoolLabels(nil)
}

func (p *NodePoolLabels) columns() []string {
	return []string{
		"ID",
		"Label",
		"Created At",
		"Updated At",
	}
}

// Table outputs the table representation to the given writer
func (p *NodePoolLabels) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, l := range p.items {
		table.Append([]string{
			string(l.ID),
			buildNodePoolLabelString(l),
			convert.UnixTimeMSToString(l.CreatedAt),
			convert.UnixTimeMSToString(l.UpdatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *NodePoolLabels) JSON(w io.Writer, listView bool) error {
	if !listView && len(p.items) == 1 {
		return displayJSON(w, p.items[0])
	}

	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *NodePoolLabels) YAML(w io.Writer, listView bool) error {
	if !listView && len(p.items) == 1 {
		return displayYAML(w, p.items[0])
	}

	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *NodePoolLabels) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}

// buildLabel string returns a string representing the given label.
// It assumes that the label is well-formed (i.e. it has a non-empty key)
func buildNodePoolLabelString(l types.NodePoolLabel) string {
	var key, value string
	if l.Key != nil {
		key = *l.Key
	}

	if l.Value != nil {
		value = *l.Value
	}

	// We're assuming the label is well-formed, so ignore errors
	str, _ := label.String(key, value)

	return str
}
