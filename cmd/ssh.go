package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud"
)

var (
	username     string
	identityFile string
	nodePoolIDs  []string

	syncPanes bool
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into nodes of a cluster",
	Long: `SSH into nodes of a cluster

Node pools to SSH into may be provided via a comma-separated list using --node-pools.

If no node pools are explicitly provided, the default is to SSH into all nodes in all pools.

tmux is required in order to split the terminal into multiple panes - one for each SSH session.

For more information on tmux, please visit https://github.com/tmux/tmux.`,

	PersistentPreRunE: clientsetRequiredPreRunE,
	PreRunE:           clusterScopedPreRunE,

	RunE: func(_ *cobra.Command, args []string) error {
		if nodePoolIDs == nil {
			// Default is to SSH into all node pools
			var err error
			nodePoolIDs, err = getAllNodePoolIDsFromCloud(clientset, clusterID)
			if err != nil {
				return err
			}
		}

		ips := make([]string, 0)
		for _, poolID := range nodePoolIDs {
			if poolID == "" {
				// This can happen due to oddities in how a user specified --node-pools
				continue
			}

			nodes, err := clientset.Provision().Nodes(organizationID, clusterID, poolID).List()
			if err != nil {
				return errors.Wrapf(err, "listing nodes for node pool %q", poolID)
			}

			for _, node := range nodes {
				ips = append(ips, node.Addresses.ExternalIP)
			}
		}

		if len(ips) == 0 {
			fmt.Println("No nodes to SSH into. Are your command line flags correct?")
			return nil
		}

		// First arg must be program name
		tmuxPath, err := exec.LookPath("tmux")
		if err != nil {
			return errors.Wrapf(err, "tmux is required but couldn't be found (is it installed?)")
		}

		tmuxArgs := []string{"tmux"}
		for i, ip := range ips {
			// Do not quote command, as it will be passed properly to tmux via
			// syscall.Exec as long as the entire command is a single string in
			// the slice of args.
			sshCmd := fmt.Sprintf("ssh %s@%s", username, ip)
			if identityFile != "" {
				sshCmd = fmt.Sprintf("%s -i %s", sshCmd, identityFile)
			}

			if i == 0 {
				tmuxArgs = append(tmuxArgs, "new-window")
			} else {
				tmuxArgs = append(tmuxArgs, "split-window")
			}

			// Semicolons are used to delimit commands to tmux
			tmuxArgs = append(tmuxArgs, fmt.Sprintf("%s;", sshCmd))
		}

		tmuxArgs = append(tmuxArgs, "select-layout", "tiled;")

		if syncPanes {
			tmuxArgs = append(tmuxArgs, "set-window-option", "synchronize-panes", "on;")
		}

		// Use syscall.Exec instead of os.Exec because we want to replace the
		// current process, not spawn an external one. This will only return
		// if there is problem executing the given binary.
		err = syscall.Exec(tmuxPath, tmuxArgs, os.Environ())
		if err != nil {
			return errors.Wrapf(err, "running command %v", tmuxArgs)
		}

		// Unreachable
		panic(nil)
	},
}

func getAllNodePoolIDsFromCloud(clientset cloud.Interface, clusterID string) ([]string, error) {
	poolIDs := make([]string, 0)

	pools, err := clientset.Provision().NodePools(organizationID, clusterID).List()
	if err != nil {
		return nil, errors.Wrap(err, "listing node pools")
	}

	for _, pool := range pools {
		poolIDs = append(poolIDs, string(pool.ID))
	}

	return poolIDs, nil
}

func init() {
	rootCmd.AddCommand(sshCmd)

	requireClientset(sshCmd)
	bindCommandToClusterScope(sshCmd, false)

	sshCmd.Flags().StringSliceVar(&nodePoolIDs, "node-pools", nil, "comma-separated node pool IDs to SSH into (default is all node pools)")
	sshCmd.Flags().StringVarP(&username, "username", "u", "containership", "username")
	sshCmd.Flags().StringVarP(&identityFile, "identity-file", "i", "", "identity file (passed to ssh -i)")
	sshCmd.Flags().BoolVarP(&syncPanes, "sync", "s", false, "synchronize tmux panes")
}
