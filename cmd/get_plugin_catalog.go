package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// Flags
var (
	pluginType           string
	pluginImplementation string
	pluginVersion        string
)

// getPluginCatalogCmd represents the getPluginCatalog command
var getPluginCatalogCmd = &cobra.Command{
	Use:     "plugin-catalog",
	Short:   "Get the plugin catalog",
	Aliases: resource.PluginCatalog().Aliases(),

	Long: `Get the plugin catalog, which is a list of all supported plugins with
version compatibility details.

To get the full plugin catalog, simply run:

	csctl get plugin-catalog

You may also wish to filter the results by plugin type, implementation, and version. For example:

	# Get all container network interface (CNI) plugins
	csctl get plugin-catalog --type cni

	# Get all versions of the Calico CNI plugin (note that --type is required)
	csctl get plugin-catalog --type cni --implementation calico

	# Get a specific version of the Calico CNI plugin (note that all flags are required)
	csctl get plugin-catalog --type cni --implementation calico --version 2.0.0`,

	RunE: func(cmd *cobra.Command, args []string) error {
		var plugins *resource.PluginsCatalog
		switch {
		case pluginType == "" && pluginImplementation == "" && pluginVersion == "":
			// Nothing specified - get the entire catalog
			pc, err := clientset.API().PluginCatalog().Get()
			if err != nil {
				return err
			}

			plugins = resource.NewPluginCatalog(pc)

		case pluginType != "" && pluginImplementation != "" && pluginVersion != "":
			// Everything specified - get the specific plugin version
			v, err := clientset.API().PluginCatalog().TypeImplementationVersion(
				pluginType, pluginImplementation, pluginVersion)
			if err != nil {
				return err
			}

			plugins = resource.NewPluginCatalogFromVersion(pluginType, pluginImplementation, v)

		case pluginType != "" && pluginImplementation != "":
			// Everything but version specified - get the plugin definition
			// for this type and implementation
			def, err := clientset.API().PluginCatalog().TypeImplementation(
				pluginType, pluginImplementation)
			if err != nil {
				return err
			}

			plugins = resource.NewPluginCatalogFromDefinition(pluginType, def)

		case pluginType != "":
			defs, err := clientset.API().PluginCatalog().Type(pluginType)
			if err != nil {
				return err
			}

			plugins = resource.NewPluginCatalogFromDefinitions(pluginType, defs)

		default:
			return errors.New("invalid flag combination - see usage")
		}

		outputResponse(plugins, false)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getPluginCatalogCmd)

	getPluginCatalogCmd.Flags().StringVarP(&pluginType, "type", "t", "", "filter by plugin type")
	getPluginCatalogCmd.Flags().StringVarP(&pluginImplementation, "implementation", "i", "", "filter by plugin implementation (type must also be provided)")
	getPluginCatalogCmd.Flags().StringVarP(&pluginVersion, "version", "v", "", "get a specific plugin version (type, implementation must also be provided)")
}
