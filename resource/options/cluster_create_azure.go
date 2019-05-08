package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

// AzureClusterCreate is the set of options required
// to create a Azure cluster
type AzureClusterCreate struct {
	ClusterCreate
}

const (
	azureCCM = "azure"
	azureCNI = "canal"
)

// DefaultAndValidate defaults and validates all options
func (o *AzureClusterCreate) DefaultAndValidate() error {
	if err := o.ClusterCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.defaultAndValidateCNI(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCNI)
	}

	if err := o.defaultAndValidateCCM(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCCM)
	}

	return nil
}

// CreateCKEClusterRequest constructs a CreateCKEClusterRequest from these options
func (o *AzureClusterCreate) CreateCKEClusterRequest() types.CreateCKEClusterRequest {
	return types.CreateCKEClusterRequest{
		ProviderID: types.UUID(o.ProviderID),
		TemplateID: types.UUID(o.TemplateID),
		Labels:     o.labels,
		Plugins:    o.plugins,
	}
}

func (o *AzureClusterCreate) defaultAndValidateCNI() error {
	impl, version, err := o.PluginCNIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return errors.Errorf("CNI plugin is required (can't specify an implementation of %q)", plugin.NoImplementation)
	}

	if impl != "" && impl != azureCNI {
		return errors.Errorf("only %s CNI implementation is allowed", azureCNI)
	}
	impl = azureCNI

	pType := "cni"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *AzureClusterCreate) defaultAndValidateCCM() error {
	impl, version, err := o.PluginCCMFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != azureCCM {
		return errors.Errorf("only %s CCM implementation is allowed", azureCCM)
	}
	impl = azureCCM

	pType := "cloud_controller_manager"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}
