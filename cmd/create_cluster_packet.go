package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
)

var packetCreateClusterOpts options.PacketClusterCreate

// createClusterPacketCmd represents the createClusterPacket command
var createClusterPacketCmd = &cobra.Command{
	Use:   "packet",
	Short: "Create a Packet cluster",
	Args:  cobra.NoArgs,
	Long: `Create a Packet cluster.

For general usage info for cluster creation, see:

    csctl create cluster --help`,

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		packetCreateClusterOpts.ClusterCreate = createClusterOpts

		if err := packetCreateClusterOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := packetCreateClusterOpts.CreateCKEClusterRequest()

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createClusterCmd.AddCommand(createClusterPacketCmd)

	bindCommandToOrganizationScope(createClusterPacketCmd, false)

	// These are defined in the parent command, but they aren't required there
	_ = createClusterPacketCmd.MarkPersistentFlagRequired("template")
	_ = createClusterPacketCmd.MarkPersistentFlagRequired("provider")
	_ = createClusterPacketCmd.MarkPersistentFlagRequired("name")
	_ = createClusterPacketCmd.MarkPersistentFlagRequired("environment")
}
