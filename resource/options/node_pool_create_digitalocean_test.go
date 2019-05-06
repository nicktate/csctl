package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDONodePoolDefaultAndValidate(t *testing.T) {
	var opts = DigitalOceanNodePoolCreate{
		NodePoolCreate: validMasterNodePoolCreateOpts,
	}
	opts.KubernetesVersion = "invalid"
	err := opts.DefaultAndValidate()
	assert.Error(t, err, "bad parent opts results in error")
}
