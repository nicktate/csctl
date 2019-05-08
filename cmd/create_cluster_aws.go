package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
)

var awsClusterOpts options.AWSClusterCreate

// createClusterAWSCmd represents the createClusterAWS command
var createClusterAWSCmd = &cobra.Command{
	Use:   "aws",
	Short: "Create an AWS cluster",
	Args:  cobra.NoArgs,
	Long: `Create an AWS cluster.

For general usage info for cluster creation, see:

    csctl create cluster --help`,

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		awsClusterOpts.ClusterCreate = createClusterOpts

		if err := awsClusterOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := awsClusterOpts.CreateCKEClusterRequest()

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createClusterCmd.AddCommand(createClusterAWSCmd)

	bindCommandToOrganizationScope(createClusterAWSCmd, false)

	// These are defined in the parent command, but they aren't required there
	_ = createClusterAWSCmd.MarkPersistentFlagRequired("template")
	_ = createClusterAWSCmd.MarkPersistentFlagRequired("provider")
	_ = createClusterAWSCmd.MarkPersistentFlagRequired("name")
	_ = createClusterAWSCmd.MarkPersistentFlagRequired("environment")
}
