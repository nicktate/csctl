package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// getNodePoolLabelCmd represents the getNodePoolLabel command
var getNodePoolLabelCmd = &cobra.Command{
	Use:     "node-pool-label",
	Short:   "Get a node pool label or list of node pool labels",
	Aliases: resource.NodePoolLabel().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.NodePoolLabel, 1)
		var err error
		if len(args) == 1 {
			var l *types.NodePoolLabel
			l, err = clientset.Provision().NodePoolLabels(organizationID, clusterID, nodePoolID).Get(args[0])
			resp[0] = *l
		} else {
			resp, err = clientset.Provision().NodePoolLabels(organizationID, clusterID, nodePoolID).List()
		}

		if err != nil {
			return err
		}

		labels := resource.NewNodePoolLabels(resp)

		outputResponse(labels)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getNodePoolLabelCmd)

	bindCommandToNodePoolScope(getNodePoolLabelCmd, false)
}
