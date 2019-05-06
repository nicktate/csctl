package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validMasterNodePoolCreateOpts = NodePoolCreate{
	Mode:              "master",
	Name:              "name",
	Count:             3,
	KubernetesVersion: "1.13.4",
	OperatingSystem:   "ubuntu",
}

func TestNodePoolCreateDefaultAndValidate(t *testing.T) {
	var opts = NodePoolCreate{}
	err := opts.DefaultAndValidate()
	assert.Error(t, err, "empty opts not ok")

	opts = validMasterNodePoolCreateOpts

	// Fields that are not user-settable but required obtain a value
	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "valid options")
	assert.NotEmpty(t, opts.nodePoolType, "nodePoolType set")

	opts.KubernetesVersion = "invalid"
	err = opts.DefaultAndValidate()
	assert.Error(t, err, "invalid semver")
}

func TestTemplateNodePool(t *testing.T) {
	opts := validMasterNodePoolCreateOpts
	err := opts.DefaultAndValidate()
	assert.NoError(t, err, "valid options")

	tmpl := opts.TemplateNodePool()
	assert.NotNil(t, tmpl)
	assert.True(t, tmpl.Etcd, "etcd enabled for masters")

	opts.Mode = "worker"
	tmpl = opts.TemplateNodePool()
	assert.NotNil(t, tmpl)
	assert.False(t, tmpl.Etcd, "etcd not enabled for workers")
}
