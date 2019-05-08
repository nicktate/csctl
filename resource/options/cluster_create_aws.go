package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

// AWSClusterCreate is the set of options required
// to create a AWS cluster
type AWSClusterCreate struct {
	ClusterCreate
}

const (
	awsCNI = "calico"
	awsCCM = "amazon_web_services"
	awsCSI = "amazon_web_services"
)

// DefaultAndValidate defaults and validates all options
func (o *AWSClusterCreate) DefaultAndValidate() error {
	if err := o.ClusterCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.defaultAndValidateCNI(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCNI)
	}

	if err := o.defaultAndValidateCSI(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCSI)
	}

	if err := o.defaultAndValidateCCM(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCCM)
	}

	return nil
}

// CreateCKEClusterRequest constructs a CreateCKEClusterRequest from these options
func (o *AWSClusterCreate) CreateCKEClusterRequest() types.CreateCKEClusterRequest {
	return types.CreateCKEClusterRequest{
		ProviderID: types.UUID(o.ProviderID),
		TemplateID: types.UUID(o.TemplateID),
		Labels:     o.labels,
		Plugins:    o.plugins,
	}
}

func (o *AWSClusterCreate) defaultAndValidateCNI() error {
	impl, version, err := o.PluginCNIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return errors.Errorf("CNI plugin is required (can't specify an implementation of %q)", plugin.NoImplementation)
	}

	if impl != "" && impl != awsCNI {
		return errors.Errorf("only %s CNI implementation is allowed", awsCNI)
	}
	impl = awsCNI

	pType := types.CreateCKEClusterPluginTypeCni
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *AWSClusterCreate) defaultAndValidateCCM() error {
	impl, version, err := o.PluginCCMFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != awsCCM {
		return errors.Errorf("only %s CCM implementation is allowed", awsCCM)
	}
	impl = awsCCM

	pType := types.CreateCKEClusterPluginTypeCloudControllerManager
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *AWSClusterCreate) defaultAndValidateCSI() error {
	impl, version, err := o.PluginCSIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != awsCSI {
		return errors.Errorf("only %s CSI implementation is allowed", awsCSI)
	}
	impl = awsCSI

	pType := types.CreateCKEClusterPluginTypeCsi
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}
