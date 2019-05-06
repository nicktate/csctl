package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// DigitalOceanNodePoolCreate is the set of options required to create a template
type DigitalOceanNodePoolCreate struct {
	NodePoolCreate

	digitalOceanDroplet
}

// DefaultAndValidate defaults and validates all options
func (o *DigitalOceanNodePoolCreate) DefaultAndValidate() error {
	if err := o.NodePoolCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.digitalOceanDroplet.defaultAndValidate(o.OperatingSystem); err != nil {
		return errors.Wrap(err, "validating droplet options")
	}

	return nil
}

// NodePoolDigitalOceanCreateRequest constructs a NodePoolCreateRequest from these options
func (o *DigitalOceanNodePoolCreate) NodePoolDigitalOceanCreateRequest() types.NodePoolDigitalOceanCreateRequest {
	return types.NodePoolDigitalOceanCreateRequest{
		Variable: o.TemplateNodePool(),
		Resource: &types.DigitalOceanDropletConfiguration{
			Image:  &o.Image,
			Region: &o.Region,
			Size:   &o.InstanceSize,

			Backups:           o.backups,
			Monitoring:        o.monitoring,
			PrivateNetworking: &o.privateNetworking,
		},
	}
}
