package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var createNodePoolOpts options.NodePoolCreate

// createNodePoolCmd represents the createNodePool command
var createNodePoolCmd = &cobra.Command{
	Use:     "node-pool",
	Short:   "Create (add) a node pool for an existing cluster",
	Aliases: resource.NodePool().Aliases(),
}

func init() {
	createCmd.AddCommand(createNodePoolCmd)

	bindCommandToClusterScope(createNodePoolCmd, false)

	// Required
	createNodePoolCmd.PersistentFlags().StringVarP(&createNodePoolOpts.Mode, "mode", "m", "", "node pool mode (master or worker)")
	_ = createNodePoolCmd.MarkPersistentFlagRequired("mode")
	createNodePoolCmd.PersistentFlags().StringVarP(&createNodePoolOpts.Name, "name", "n", "", "name for this node pool")
	_ = createNodePoolCmd.MarkPersistentFlagRequired("name")
	createNodePoolCmd.PersistentFlags().StringVar(&createNodePoolOpts.KubernetesVersion, "kubernetes-version", "", "Kubernetes version for the node pool")
	_ = createNodePoolCmd.MarkPersistentFlagRequired("kubernetes-version")
	createNodePoolCmd.PersistentFlags().StringVar(&createNodePoolOpts.OperatingSystem, "os", "", "Operating system (ubuntu or centos)")
	_ = createNodePoolCmd.MarkPersistentFlagRequired("os")

	// Defaulted (either here or by cloud)
	createNodePoolCmd.PersistentFlags().Int32VarP(&createNodePoolOpts.Count, "count", "c", 1, "number of nodes in the node pool (default 1)")
	createNodePoolCmd.PersistentFlags().StringVar(&createNodePoolOpts.DockerVersion, "docker-version", "", "Docker version for the node pool (defaulted according to Kubernetes version)")
}
