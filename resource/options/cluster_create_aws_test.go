package options

import (
	"testing"

	"github.com/containership/csctl/resource/plugin"
	"github.com/stretchr/testify/assert"
)

var (
	validAWSClusterCreateOpts = ClusterCreate{
		TemplateID:  validUUID,
		ProviderID:  validUUID,
		Name:        "name",
		Environment: "env",
	}
)

func TestAWSClusterCreateDefaultAndValidate(t *testing.T) {
	emptyOpts := AWSClusterCreate{}
	err := emptyOpts.DefaultAndValidate()
	assert.Error(t, err, "empty parent opts is not ok")

	opts := AWSClusterCreate{ClusterCreate: validAWSClusterCreateOpts}

	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "empty DO opts is ok")
}

func TestAWSCreateCKEClusterRequest(t *testing.T) {
	emptyOpts := AWSClusterCreate{}
	req := emptyOpts.CreateCKEClusterRequest()
	assert.NotNil(t, req, "CKE cluster request is never nil")
}

func TestAWSDefaultAndValidateCNI(t *testing.T) {
	opts := AWSClusterCreate{ClusterCreate: validAWSClusterCreateOpts}

	err := opts.defaultAndValidateCNI()
	assert.NoError(t, err, "empty CNI flag is ok")

	opts.PluginCNIFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "invalid CNI plugin flag")

	opts.PluginCNIFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "disabling CNI is not allowed")
}

func TestAWSDefaultAndValidateCCM(t *testing.T) {
	opts := AWSClusterCreate{ClusterCreate: validAWSClusterCreateOpts}

	err := opts.defaultAndValidateCCM()
	assert.NoError(t, err, "empty CCM flag is ok")

	opts.PluginCCMFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCCM()
	assert.Error(t, err, "invalid CCM plugin flag")

	opts.PluginCCMFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCCM()
	assert.NoError(t, err, "disabling CCM is allowed")
}

func TestAWSDefaultAndValidateCSI(t *testing.T) {
	opts := AWSClusterCreate{ClusterCreate: validAWSClusterCreateOpts}

	err := opts.defaultAndValidateCSI()
	assert.NoError(t, err, "empty CSI flag is ok")

	opts.PluginCSIFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCSI()
	assert.Error(t, err, "invalid CSI plugin flag")

	opts.PluginCSIFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCSI()
	assert.NoError(t, err, "disabling CSI is allowed")
}
