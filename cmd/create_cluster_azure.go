package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
)

var azureCreateClusterOpts options.AzureClusterCreate

// createClusterAzureCmd represents the createClusterAzure command
var createClusterAzureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Create an Azure cluster",
	Args:  cobra.NoArgs,
	Long: `Create an Azure cluster.

For general usage info for cluster creation, see:

    csctl create cluster --help`,

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		azureCreateClusterOpts.ClusterCreate = createClusterOpts

		if err := azureCreateClusterOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := azureCreateClusterOpts.CreateCKEClusterRequest()

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createClusterCmd.AddCommand(createClusterAzureCmd)

	bindCommandToOrganizationScope(createClusterAzureCmd, false)

	// These are defined in the parent command, but they aren't required there
	_ = createClusterAzureCmd.MarkPersistentFlagRequired("template")
	_ = createClusterAzureCmd.MarkPersistentFlagRequired("provider")
	_ = createClusterAzureCmd.MarkPersistentFlagRequired("name")
	_ = createClusterAzureCmd.MarkPersistentFlagRequired("environment")
}
