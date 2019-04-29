package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/options"
	"github.com/containership/csctl/resource/plugin"
)

var createClusterOpts options.ClusterCreate

// createClusterCmd represents the createCluster command
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create (provision) a CKE cluster",
	Long: `Create (provision) a Containership Kubernetes Engine (CKE) cluster.

If creating a cluster from a file, do not specify a provider to use and instead simply run:

csctl create cluster -f <filename>

This file must be json (TODO support yaml). All other cluster-specific flags will be ignored if -f is used.

Otherwise, please use a provider subcommand to create a cluster with more advanced, provider-specific flags.

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

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open(filename)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}
		defer f.Close()

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			return errors.Wrap(err, "reading file")
		}

		var req types.CreateCKEClusterRequest

		err = json.Unmarshal(bytes, &req)
		if err != nil {
			return errors.Wrap(err, "unmarshalling file into request type")
		}

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createCmd.AddCommand(createClusterCmd)

	createClusterCmd.Flags().StringVarP(&filename, "filename", "f", "", "create a cluster from the given file (TODO must be json for now)")
	createClusterCmd.MarkFlagRequired("filename")

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.TemplateID, "template", "t", "", "template ID to create from")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.ProviderID, "provider", "p", "", "provider ID (credentials) to use for provisioning")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Name, "name", "n", "", "cluster name")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Environment, "environment", "e", "", "environment")

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
