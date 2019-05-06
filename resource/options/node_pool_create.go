package options

import (
	"github.com/Masterminds/semver"
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// NodePoolCreate is the set of options required to create a template
type NodePoolCreate struct {
	// Required
	Mode              string
	Name              string
	Count             int32
	KubernetesVersion string
	OperatingSystem   string

	// Defaultable by cloud
	DockerVersion string

	// Not user-settable; always defaulted
	nodePoolType string
}

// DefaultAndValidate defaults and validates all options
func (o *NodePoolCreate) DefaultAndValidate() error {
	if err := o.validateKubernetesVersion(); err != nil {
		return errors.Wrap(err, "kubernetes version")
	}

	// Note that Docker versions are not validated; we delegate that to cloud
	// for simplicity

	o.nodePoolType = "node_pool"

	return nil
}

// TemplateNodePool constructs the variable block for these options
func (o *NodePoolCreate) TemplateNodePool() *types.TemplateNodePool {
	// All master node pools must run etcd
	etcdEnabled := o.Mode == "master"

	return &types.TemplateNodePool{
		Name:              &o.Name,
		Count:             &o.Count,
		KubernetesMode:    &o.Mode,
		KubernetesVersion: &o.KubernetesVersion,
		Os:                &o.OperatingSystem,
		DockerVersion:     o.DockerVersion,
		Type:              &o.nodePoolType,
		Etcd:              etcdEnabled,
	}
}

func (o *NodePoolCreate) validateKubernetesVersion() error {
	mv, err := semver.NewVersion(o.KubernetesVersion)
	if err != nil {
		return errors.Wrap(err, "validating semver")
	}
	// Note that String() returns the version with the leading 'v' stripped
	// if applicable, which is what we want for cloud interactions.
	o.KubernetesVersion = mv.String()

	return nil
}
