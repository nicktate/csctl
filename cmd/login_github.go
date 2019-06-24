package cmd

import (
	"github.com/spf13/cobra"
	"github.com/containership/csctl/pkg/oauth"
	"os/exec"
	"time"
)

// getNodeCmd represents the getNode command
var loginGithubCmd = &cobra.Command{
	Use:     "github",
	Short:   "Login with github",
	// Aliases: ,

	Args: cobra.MaximumNArgs(1),

	// PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		oauth.StartServer()

		// handle err
		openBrowserCmd := exec.Command("open", "http://localhost:8080")
		openBrowserCmd.Run()

		timer := time.NewTimer(60 * time.Second)
		<- timer.C

		oauth.StopServer()
		return nil
	},
}

func init() {
	loginCmd.AddCommand(loginGithubCmd)
}
