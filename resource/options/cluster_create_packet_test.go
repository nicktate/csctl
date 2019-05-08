package options

import (
	"testing"

	"github.com/containership/csctl/resource/plugin"
	"github.com/stretchr/testify/assert"
)

var (
	validPacketClusterCreateOpts = ClusterCreate{
		TemplateID:  validUUID,
		ProviderID:  validUUID,
		Name:        "name",
		Environment: "env",
	}
)

func TestPacketClusterCreateDefaultAndValidate(t *testing.T) {
	emptyOpts := PacketClusterCreate{}
	err := emptyOpts.DefaultAndValidate()
	assert.Error(t, err, "empty parent opts is not ok")

	opts := PacketClusterCreate{ClusterCreate: validPacketClusterCreateOpts}

	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "empty DO opts is ok")
}

func TestPacketCreateCKEClusterRequest(t *testing.T) {
	emptyOpts := PacketClusterCreate{}
	req := emptyOpts.CreateCKEClusterRequest()
	assert.NotNil(t, req, "CKE cluster request is never nil")
}

func TestPacketDefaultAndValidateCNI(t *testing.T) {
	opts := PacketClusterCreate{ClusterCreate: validPacketClusterCreateOpts}

	err := opts.defaultAndValidateCNI()
	assert.NoError(t, err, "empty CNI flag is ok")

	opts.PluginCNIFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "invalid CNI plugin flag")

	opts.PluginCNIFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "disabling CNI is not allowed")
}

func TestPacketDefaultAndValidateCCM(t *testing.T) {
	opts := PacketClusterCreate{ClusterCreate: validPacketClusterCreateOpts}

	err := opts.defaultAndValidateCCM()
	assert.NoError(t, err, "empty CCM flag is ok")

	opts.PluginCCMFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCCM()
	assert.Error(t, err, "invalid CCM plugin flag")

	opts.PluginCCMFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCCM()
	assert.NoError(t, err, "disabling CCM is allowed")
}
