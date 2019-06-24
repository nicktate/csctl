package cmd

import (
	"github.com/spf13/cobra"
)

// Flags
var (
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Containership Cloud",

	PersistentPreRunE: clientsetRequiredPreRunE,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	requireClientset(loginCmd)
}
