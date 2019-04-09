package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/label"
	"github.com/containership/csctl/resource/table"
)

// ClusterLabels is a list of the associated cloud resource with additional functionality
type ClusterLabels struct {
	resource
	items []types.ClusterLabel
}

// NewClusterLabels constructs a new ClusterLabels wrapping the given cloud type
func NewClusterLabels(items []types.ClusterLabel) *ClusterLabels {
	return &ClusterLabels{
		resource: resource{
			name:    "cluster-label",
			plural:  "cluster-labels",
			aliases: []string{"cl", "cls"},
		},
		items: items,
	}
}

// ClusterLabel constructs a new ClusterLabels with no underlying items, useful for
// interacting with the metadata itself.
func ClusterLabel() *ClusterLabels {
	return NewClusterLabels(nil)
}

func (p *ClusterLabels) columns() []string {
	return []string{
		"ID",
		"Label",
		"Created At",
		"Updated At",
	}
}

// Table outputs the table representation to the given writer
func (p *ClusterLabels) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, l := range p.items {
		table.Append([]string{
			string(l.ID),
			buildClusterLabelString(l),
			convert.UnixTimeMSToString(l.CreatedAt),
			convert.UnixTimeMSToString(l.UpdatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *ClusterLabels) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *ClusterLabels) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *ClusterLabels) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}

// buildLabel string returns a string representing the given label.
// It assumes that the label is well-formed (i.e. it has a non-empty key)
func buildClusterLabelString(l types.ClusterLabel) string {
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
