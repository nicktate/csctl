package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
	"github.com/containership/csctl/resource/plugin"
)

var createClusterOpts options.ClusterCreate

// createClusterCmd represents the createCluster command
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create (provision) a CKE cluster",
	Long: `Create (provision) a Containership Kubernetes Engine (CKE) cluster.

Each cloud provider has its own subcommand. Generic flags belong to this root command.

Clusters are created from a template which is specified using --template (-t).
For more information on templates, see the template help:

	csctl get template --help

The cloud provider credentials to use are also required, and can be specified
using --provider (-p). To add new cloud provider credentials, visit the
Organization Settings page in the Containership UI.

To list available providers, use:

	csctl get providers

Plugins are specified using flags corresponding to the plugin type. For
example, to specify the autoscaler plugin, use --plugin-autoscaler. The
implementation and version are provided as follows, for example:

	csctl create cluster digitalocean --plugin-autoscaler=cerebral@0.3.2-alpha ...

The version can also be omitted, in which case the latest compatible version is used.

Many best-practice plugins are added by default, meaning that plugins only have
to be explicitly defined if you want fine-grained control over what is added to
your newly provisioned cluster.

To view all available plugins, use the plugin catalog:

	csctl get plugin-catalog`,
}

func init() {
	createCmd.AddCommand(createClusterCmd)

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.TemplateID, "template", "t", "", "template ID to create from")
	createClusterCmd.MarkPersistentFlagRequired("template")

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.ProviderID, "provider", "p", "", "provider ID (credentials) to use for provisioning")
	createClusterCmd.MarkPersistentFlagRequired("provider")

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Name, "name", "n", "", "cluster name")
	createClusterCmd.MarkPersistentFlagRequired("name")

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Environment, "environment", "e", "", "environment")
	createClusterCmd.MarkPersistentFlagRequired("environment")

	// Plugins
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginCNIFlag.Val, "plugin-cni", "",
		"Container Networking Interface (CNI) plugin")
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginCSIFlag.Val, "plugin-csi", "",
		fmt.Sprintf("Cloud Storage Interface (CSI) plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginCCMFlag.Val, "plugin-ccm", "",
		fmt.Sprintf("Cloud Controller Manager (CCM) plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginMetricsFlag.Val, "plugin-metrics", "",
		fmt.Sprintf("metrics plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginClusterManagementFlag.Val, "plugin-cluster-management", "",
		"Cluster Management plugin (implementation must be \"containership\")")
	createClusterCmd.PersistentFlags().StringVar(&createClusterOpts.PluginAutoscalerFlag.Val, "plugin-autoscaler", "",
		"autoscaler plugin")
}
