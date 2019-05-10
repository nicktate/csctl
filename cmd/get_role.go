package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/resource"
)

// getRoleCmd represents the getRole command
var getRoleCmd = &cobra.Command{
	Use:     "role",
	Short:   "Get an authorization role or list of roles",
	Aliases: resource.AuthorizationRole().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.AuthorizationRole, 1)
		var err error
		if len(args) == 1 {
			var v *types.AuthorizationRole
			v, err = clientset.Auth().AuthorizationRoles(organizationID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Auth().AuthorizationRoles(organizationID).List()
		}

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
	getCmd.AddCommand(getRoleCmd)

	bindCommandToOrganizationScope(getRoleCmd, false)
}
