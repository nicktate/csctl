package cmd

import (
	"fmt"
	"os"

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

By default, the Kubeconfig is printed to stdout. To output to a file instead, use --filename (-f).

For example:

	# Export Kubeconfig for a CKE cluster
	csctl export kubeconfig --cluster <cluster_id> --filename admin.conf

	# Interact with cluster using kubectl
	kubectl --kubeconfig admin.conf get pods --all-namespaces
`,

	Args: cobra.NoArgs,

	PersistentPreRunE: clientsetRequiredPreRunE,
	PreRunE:           clusterScopedPreRunE,

	Run: func(cmd *cobra.Command, args []string) {
		if organizationID == "" || clusterID == "" {
			fmt.Println("organization and cluster are required")
			return
		}

		// TODO do this better once proxy client is in place; see issue #7
		proxyBaseURL := viper.GetString("proxyBaseURL")
		serverAddress := fmt.Sprintf("%s/v3/organizations/%s/clusters/%s/k8sapi/proxy",
			proxyBaseURL, organizationID, clusterID)

		account, err := clientset.API().Account().Get()
		if err != nil {
			fmt.Println(err)
			return
		}

		cluster, err := clientset.API().Clusters(organizationID).Get(clusterID)
		if err != nil {
			fmt.Println(err)
			return
		}

		// TODO error handling
		// TODO UUID typecasting
		cfg := kubeconfig.New(&kubeconfig.Config{
			ServerAddress: serverAddress,
			ClusterID:     string(cluster.ID),
			UserID:        string(account.ID),
			Token:         userToken,
		})

		w := os.Stdout
		if filename != "" {
			w, err = os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
			defer w.Close()
		}

		// TODO implement merging into ~/.kube/config, which should be the new default
		// (instead of stdout)
		err = kubeconfig.Write(cfg, w)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	exportCmd.AddCommand(exportKubeconfigCmd)

	requireClientset(exportKubeconfigCmd)
	bindCommandToClusterScope(exportKubeconfigCmd, false)

	exportKubeconfigCmd.Flags().StringVarP(&filename, "filename", "f", "", "output kubeconfig to file (default is stdout)")
}
