package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

// GoogleClusterCreate is the set of options required
// to create a Google cluster
type GoogleClusterCreate struct {
	ClusterCreate
}

const (
	googleCNI = "calico"
	googleCCM = "google_cloud_platform"
	googleCSI = "google_cloud_platform"
)

// DefaultAndValidate defaults and validates all options
func (o *GoogleClusterCreate) DefaultAndValidate() error {
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
func (o *GoogleClusterCreate) CreateCKEClusterRequest() types.CreateCKEClusterRequest {
	return types.CreateCKEClusterRequest{
		ProviderID: types.UUID(o.ProviderID),
		TemplateID: types.UUID(o.TemplateID),
		Labels:     o.labels,
		Plugins:    o.plugins,
	}
}

func (o *GoogleClusterCreate) defaultAndValidateCNI() error {
	impl, version, err := o.PluginCNIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return errors.Errorf("CNI plugin is required (can't specify an implementation of %q)", plugin.NoImplementation)
	}

	if impl != "" && impl != googleCNI {
		return errors.Errorf("only %s CNI implementation is allowed", googleCNI)
	}
	impl = googleCNI

	pType := types.CreateCKEClusterPluginTypeCni
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *GoogleClusterCreate) defaultAndValidateCCM() error {
	impl, version, err := o.PluginCCMFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != googleCCM {
		return errors.Errorf("only %s CCM implementation is allowed", googleCCM)
	}
	impl = googleCCM

	pType := types.CreateCKEClusterPluginTypeCloudControllerManager
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *GoogleClusterCreate) defaultAndValidateCSI() error {
	impl, version, err := o.PluginCSIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != googleCSI {
		return errors.Errorf("only %s CSI implementation is allowed", googleCSI)
	}
	impl = googleCSI

	pType := types.CreateCKEClusterPluginTypeCsi
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}
