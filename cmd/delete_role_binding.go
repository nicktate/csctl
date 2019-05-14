package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteRoleBindingCmd represents the deleteRoleBinding command
var deleteRoleBindingCmd = &cobra.Command{
	Use:     "role-binding",
	Short:   "Delete one or more authorization role bindings",
	Aliases: resource.AuthorizationRoleBinding().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Auth().AuthorizationRoleBindings(organizationID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Authorization role %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteRoleBindingCmd)

	bindCommandToOrganizationScope(deleteRoleBindingCmd, false)
}
