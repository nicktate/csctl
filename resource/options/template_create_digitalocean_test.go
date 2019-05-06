package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDODefaultAndValidate(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{}
	// Everything is defaultable, so no error should occur for empty opts
	err := opts.DefaultAndValidate()
	assert.Nil(t, err, "empty opts ok")

	// Fields that are not user-settable but required obtain a value
	assert.NotEmpty(t, opts.providerName, "provider name set")
}

func TestDOTemplate(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{
		TemplateCreate: TemplateCreate{
			OperatingSystem: "ubuntu",
		},
		digitalOceanDroplet: digitalOceanDroplet{
			Image:        "ubuntu-16-04-x64",
			Region:       "nyc2",
			InstanceSize: "s-2vcpu-2gb",
		},
	}
	err := opts.DefaultAndValidate()
	assert.NoError(t, err)

	req := opts.CreateTemplateRequest()
	assert.Nil(t, req.Validate(nil), "valid request created")
}

func TestDefaultAndValidateDigitalOceanDroplet(t *testing.T) {
	opts := getBaseOptions(t)

	err := opts.digitalOceanDroplet.defaultAndValidateImage(opts.OperatingSystem)
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.Image, "image set")
}

func TestDefaultAndValidateRegion(t *testing.T) {
	opts := getBaseOptions(t)

	err := opts.digitalOceanDroplet.defaultAndValidateRegion()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.Region, "region set")
}

func TestDefaultAndValidateInstanceSize(t *testing.T) {
	opts := getBaseOptions(t)

	err := opts.digitalOceanDroplet.defaultAndValidateInstanceSize()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.InstanceSize, "instance size set")
}

// get a new DO options struct with parent options already defaulted
func getBaseOptions(t *testing.T) *DigitalOceanTemplateCreate {
	var opts = DigitalOceanTemplateCreate{}
	err := opts.TemplateCreate.DefaultAndValidate()
	assert.NoError(t, err, "default and validate parent options")
	return &opts
}
