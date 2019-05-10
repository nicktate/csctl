package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/resource"
)

// Flags
var (
	roleIDForBindings string
)

// getRoleBindingCmd represents the getRoleBinding command
var getRoleBindingCmd = &cobra.Command{
	Use:     "role-binding",
	Short:   "Get an authorization role binding or list of role bindings",
	Aliases: resource.AuthorizationRoleBinding().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.AuthorizationRoleBinding, 1)
		var err error
		if len(args) == 1 {
			var v *types.AuthorizationRoleBinding
			v, err = clientset.Auth().AuthorizationRoleBindings(organizationID).Get(args[0])
			resp[0] = *v
		} else if roleIDForBindings != "" {
			resp, err = clientset.Auth().AuthorizationRoleBindings(organizationID).ListForRole(roleIDForBindings)
		} else {
			resp, err = clientset.Auth().AuthorizationRoleBindings(organizationID).List()
		}

		if err != nil {
			return err
		}

		roles := resource.NewAuthorizationRoleBindings(resp)

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
	getCmd.AddCommand(getRoleBindingCmd)

	bindCommandToOrganizationScope(getRoleBindingCmd, false)

	getRoleBindingCmd.Flags().StringVar(&roleIDForBindings, "role", "", "filter by role")
}
