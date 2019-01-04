package cmd

import (
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to containership cloud",

	// allow login commands to be run without token
	PersistentPreRunE: rootPreRunE(false),
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
