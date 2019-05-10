package cmd

import (
	"github.com/spf13/cobra"

	//"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/resource"
)

// getClusterRBACCmd represents the getClusterRBAC command
var getClusterRBACCmd = &cobra.Command{
	Use:   "cluster-rbac",
	Short: "Get all Kubernetes RBAC rules for a given cluster",

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := clientset.Auth().AuthorizationRoles(organizationID).KubernetesClusterRBAC(clusterID)
		if err != nil {
			return err
		}

		roles := resource.NewAuthorizationRoles(resp)

		if mineOnly {
			me, err := getMyAccountID()
			if err != nil {
				return err
			}

			roles.FilterByOwnerID(me)
		}

		outputResponse(roles, len(args) != 1)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterRBACCmd)

	bindCommandToClusterScope(getClusterRBACCmd, false)
}
