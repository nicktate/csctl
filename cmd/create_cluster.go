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
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginCNIFlag.Val, "plugin-cni", "",
		"Container Networking Interface (CNI) plugin")
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginCSIFlag.Val, "plugin-csi", "",
		fmt.Sprintf("Cloud Storage Interface (CSI) plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginCCMFlag.Val, "plugin-ccm", "",
		fmt.Sprintf("Cloud Controller Manager (CCM) plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginMetricsFlag.Val, "plugin-metrics", "",
		fmt.Sprintf("metrics plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginClusterManagementFlag.Val, "plugin-cluster-management", "",
		"Cluster Management plugin (implementation must be \"containership\")")
	createClusterCmd.Flags().StringVar(&createClusterOpts.PluginAutoscalerFlag.Val, "plugin-autoscaler", "",
		"autoscaler plugin")
}
