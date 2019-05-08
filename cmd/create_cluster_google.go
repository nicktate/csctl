package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
)

var googleCreateClusterOpts options.GoogleClusterCreate

// createClusterGoogleCmd represents the createClusterGoogle command
var createClusterGoogleCmd = &cobra.Command{
	Use:     "google",
	Short:   "Create a Google (GCP) cluster",
	Args:    cobra.NoArgs,
	Aliases: []string{"gcp"},
	Long: `Create a Google (GCP) cluster.

For general usage info for cluster creation, see:

    csctl create cluster --help`,

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		googleCreateClusterOpts.ClusterCreate = createClusterOpts

		if err := googleCreateClusterOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := googleCreateClusterOpts.CreateCKEClusterRequest()

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createClusterCmd.AddCommand(createClusterGoogleCmd)

	bindCommandToOrganizationScope(createClusterGoogleCmd, false)

	// These are defined in the parent command, but they aren't required there
	_ = createClusterGoogleCmd.MarkPersistentFlagRequired("template")
	_ = createClusterGoogleCmd.MarkPersistentFlagRequired("provider")
	_ = createClusterGoogleCmd.MarkPersistentFlagRequired("name")
	_ = createClusterGoogleCmd.MarkPersistentFlagRequired("environment")
}
