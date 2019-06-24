package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// TODO viper dependency should not be needed here
	"github.com/spf13/viper"

	"github.com/containership/csctl/pkg/kubeconfig"
)

// Flags
var (
	filename string
)

// exportKubeconfigCmd represents the export command
var exportKubeconfigCmd = &cobra.Command{
	Use:     "kubeconfig",
	Short:   "Export a kubeconfig for a cluster",
	Aliases: []string{"kubecfg"},
	Long: `Export a Kubeconfig for interacting with a cluster via e.g. kubectl

By default, the Kubeconfig is merged into the config specified by the KUBECONFIG environment variable.
If KUBECONFIG is not set, it defaults to ~/.kube/config.
This behavior is identical to that of kubectl.

Example using merged default config:

	# Export Kubeconfig for a CKE cluster
	csctl export kubeconfig --cluster <cluster_id>

	# Interact with cluster using kubectl
	kubectl get pods --all-namespaces

To output to a file instead, use --filename (-f).
In this case, the KUBECONFIG environment variable is ignored and merging is _not_ performed.
This will simply write a new config file with only the configuration for the specified cluster.

Example of using a file:

	# Export Kubeconfig for a CKE cluster
	csctl export kubeconfig --cluster <cluster_id> --filename admin.conf

	# Interact with cluster using kubectl
	kubectl --kubeconfig admin.conf get pods --all-namespaces
`,

	Args: cobra.NoArgs,

	PersistentPreRunE: clientsetRequiredPreRunE,
	PreRunE:           clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO do this better once proxy client is in place; see issue #7
		proxyBaseURL := viper.GetString("proxyBaseURL")
		serverAddress := fmt.Sprintf("%s/v3/organizations/%s/clusters/%s/k8sapi/proxy",
			proxyBaseURL, organizationID, clusterID)

		account, err := clientset.API().Account().Get()
		if err != nil {
			return err
		}

		cluster, err := clientset.API().Clusters(organizationID).Get(clusterID)
		if err != nil {
			return err
		}

		cfg := kubeconfig.Config{
			ServerAddress: serverAddress,
			ClusterID:     string(cluster.ID),
			UserID:        string(account.ID),
			Token:         userToken,
		}

		if filename != "" {
			err = kubeconfig.WriteToFile(cfg, filename)
		} else {
			err = kubeconfig.WriteMergedDefaultConfig(cfg)
		}

		return err
	},
}

func init() {
	exportCmd.AddCommand(exportKubeconfigCmd)

	requireClientset(exportKubeconfigCmd)
	bindCommandToClusterScope(exportKubeconfigCmd, false)

	exportKubeconfigCmd.Flags().StringVarP(&filename, "filename", "f", "", "output kubeconfig to file instead of merging into file specified by KUBECONFIG")
}
