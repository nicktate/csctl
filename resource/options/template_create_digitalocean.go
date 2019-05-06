package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// DigitalOceanTemplateCreate is the set of options required
// to create a DigitalOcean template
type DigitalOceanTemplateCreate struct {
	TemplateCreate

	digitalOceanDroplet

	// Not user-settable; always defaulted
	providerName string
}

type digitalOceanDroplet struct {
	// Defaultable
	Image        string
	Region       string
	InstanceSize string

	// Not user-settable; always defaulted
	backups           bool
	monitoring        bool
	privateNetworking bool
}

const (
	digitalOceanDefaultRegion       = "nyc1"
	digitalOceanDefaultInstanceSize = "s-2vcpu-2gb"
	digitalOceanDefaultUbuntuImage  = "ubuntu-16-04-x64"
	digitalOceanDefaultCentOSImage  = "centos-7-x64"
)

// DefaultAndValidate defaults and validates all options
func (o *DigitalOceanTemplateCreate) DefaultAndValidate() error {
	if err := o.TemplateCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.digitalOceanDroplet.defaultAndValidate(o.OperatingSystem); err != nil {
		return errors.Wrap(err, "validating droplet options")
	}

	o.providerName = "digital_ocean"

	return nil
}

// CreateTemplateRequest constructs a CreateTemplateRequest from these options
func (o *DigitalOceanTemplateCreate) CreateTemplateRequest() types.CreateTemplateRequest {
	return types.CreateTemplateRequest{
		ProviderName: &o.providerName,
		Description:  &o.Description,
		Engine:       &o.engine,

		Configuration: &types.TemplateConfiguration{
			Resource: &types.TemplateResource{
				DigitaloceanDroplet: types.DigitalOceanDropletMap{
					// TODO master and worker different
					o.MasterNodePoolName: o.digitalOceanDropletConfiguration(),
					o.WorkerNodePoolName: o.digitalOceanDropletConfiguration(),
				},
			},

			Variable: o.NodePoolVariableMap(),
		},
	}
}

func (o *DigitalOceanTemplateCreate) digitalOceanDropletConfiguration() types.DigitalOceanDropletConfiguration {
	return types.DigitalOceanDropletConfiguration{
		Image:  &o.Image,
		Region: &o.Region,
		Size:   &o.InstanceSize,

		Backups:           o.backups,
		Monitoring:        o.monitoring,
		PrivateNetworking: &o.privateNetworking,
	}
}

func (o *digitalOceanDroplet) defaultAndValidate(os string) error {
	if err := o.defaultAndValidateImage(os); err != nil {
		return errors.Wrap(err, "validating image name")
	}

	if err := o.defaultAndValidateRegion(); err != nil {
		return errors.Wrap(err, "validating region")
	}

	if err := o.defaultAndValidateInstanceSize(); err != nil {
		return errors.Wrap(err, "validating instance size")
	}

	o.backups = false
	o.monitoring = false
	o.privateNetworking = true

	return nil
}

func (o *digitalOceanDroplet) defaultAndValidateImage(os string) error {
	if o.Image == "" {
		switch os {
		case "ubuntu":
			o.Image = digitalOceanDefaultUbuntuImage
		case "centos":
			o.Image = digitalOceanDefaultCentOSImage
		}
	}

	return nil
}

func (o *digitalOceanDroplet) defaultAndValidateRegion() error {
	if o.Region == "" {
		o.Region = digitalOceanDefaultRegion
	}

	return nil
}

func (o *digitalOceanDroplet) defaultAndValidateInstanceSize() error {
	if o.InstanceSize == "" {
		o.InstanceSize = digitalOceanDefaultInstanceSize
	}

	return nil
}
