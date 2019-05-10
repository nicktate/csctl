package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/resource"
)

// Flags
var (
	roleIDForRules string
)

// getRuleCmd represents the getRule command
var getRuleCmd = &cobra.Command{
	Use:     "rule",
	Short:   "Get an authorization rule or list of rules",
	Aliases: resource.AuthorizationRule().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.AuthorizationRule, 1)
		var err error
		if len(args) == 1 {
			var v *types.AuthorizationRule
			v, err = clientset.Auth().AuthorizationRules(organizationID).Get(args[0])
			resp[0] = *v
		} else if roleIDForRules != "" {
			resp, err = clientset.Auth().AuthorizationRules(organizationID).ListForRole(roleIDForRules)
		} else {
			resp, err = clientset.Auth().AuthorizationRules(organizationID).List()
		}

		if err != nil {
			return err
		}

		roles := resource.NewAuthorizationRules(resp)

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
	getCmd.AddCommand(getRuleCmd)

	bindCommandToOrganizationScope(getRuleCmd, false)

	getRuleCmd.Flags().StringVar(&roleIDForRules, "role", "", "filter by role")
}
