package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getClusterLabelCmd represents the getClusterLabel command
var getClusterLabelCmd = &cobra.Command{
	Use:     "cluster-label",
	Short:   "Get a cluster label or list of cluster labels",
	Aliases: resource.ClusterLabel().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.ClusterLabel, 1)
		var err error
		if len(args) == 1 {
			var l *types.ClusterLabel
			l, err = clientset.API().ClusterLabels(organizationID, clusterID).Get(args[0])
			resp[0] = *l
		} else {
			resp, err = clientset.API().ClusterLabels(organizationID, clusterID).List()
		}

		if err != nil {
			return err
		}

		labels := resource.NewClusterLabels(resp)

		outputResponse(labels, len(args) != 1)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterLabelCmd)

	bindCommandToClusterScope(getClusterLabelCmd, false)
}
