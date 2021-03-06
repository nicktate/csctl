package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var accessTokenOpts options.AccessTokenCreate

// createAccessTokenCmd represents the createAccessToken command
var createAccessTokenCmd = &cobra.Command{
	Use:     "access-token",
	Short:   "Create a personal access token",
	Aliases: resource.AccessToken().Aliases(),
	Long: `Create a personal access token (PAT) for authenticating with Containership Cloud.

If token creation is successful, the value will be printed to stdout. You *must* copy
the value when it is displayed, as it will not be displayed again.

The generated token can be used as the token for csctl itself, if desired.

Personal access tokens do not expire but can be deleted using:

	csctl delete access-token <access_token_id>"

For more information on PATs, please see https://docs.containership.io/developer-resources/personal-access-tokens.
`,

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := accessTokenOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := accessTokenOpts.CreateAccessTokenRequest()

		accessToken, err := clientset.API().AccessTokens().Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Access token %s created successfully. Copy the token below as it will not be displayed again:\n%s\n",
			*accessToken.Name,
			accessToken.Token)

		return nil
	},
}

func init() {
	createCmd.AddCommand(createAccessTokenCmd)
	createAccessTokenCmd.PersistentFlags().StringVarP(&accessTokenOpts.Name, "name", "n", "", "access token name")
}
