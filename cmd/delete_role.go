package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteRoleCmd represents the deleteRole command
var deleteRoleCmd = &cobra.Command{
	Use:     "role",
	Short:   "Delete one or more authorization roles",
	Aliases: resource.AuthorizationRole().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Auth().AuthorizationRoles(organizationID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Authorization role %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteRoleCmd)

	bindCommandToOrganizationScope(deleteRoleCmd, false)
}
