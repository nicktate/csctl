package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var createNodePoolDigitalOceanOpts options.DigitalOceanNodePoolCreate

// createNodePoolDigitalOceanCmd represents the createNodePoolDigitalOcean command
var createNodePoolDigitalOceanCmd = &cobra.Command{
	Use:     "digitalocean",
	Short:   "Create (add) a node pool for an existing DigitalOcean cluster",
	Args:    cobra.NoArgs,
	Aliases: []string{"do"},

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		createNodePoolDigitalOceanOpts.NodePoolCreate = createNodePoolOpts

		if err := createNodePoolDigitalOceanOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := createNodePoolDigitalOceanOpts.NodePoolDigitalOceanCreateRequest()

		np, err := clientset.Provision().NodePools(organizationID, clusterID).Create(&req)
		if err != nil {
			return err
		}

		t := resource.NewNodePools([]types.NodePool{*np})
		return t.Table(os.Stdout)
	},
}

func init() {
	createNodePoolCmd.AddCommand(createNodePoolDigitalOceanCmd)

	bindCommandToClusterScope(createNodePoolDigitalOceanCmd, false)

	createNodePoolDigitalOceanCmd.Flags().StringVar(&createNodePoolDigitalOceanOpts.Image, "image", "", "droplet image")
	createNodePoolDigitalOceanCmd.Flags().StringVar(&createNodePoolDigitalOceanOpts.Region, "region", "", "region")
	createNodePoolDigitalOceanCmd.Flags().StringVar(&createNodePoolDigitalOceanOpts.InstanceSize, "size", "", "instance size")
}
