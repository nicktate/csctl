package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of the CLI",

	Args: cobra.NoArgs,

	// allow logout command to be run without token
	PersistentPreRunE: rootPreRunE(false),

	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set("token", "")

		err := viper.WriteConfig()

		if err != nil {
			return err
		}

		fmt.Print("Successfully logged out!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
