package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Config represents the minimum required fields to construct a new kubeconfig
// for accessing Containership clusters
type Config struct {
	ServerAddress string
	ClusterID     string
	UserID        string
	Token         string
}

// WriteMergedDefaultConfig merges the new config into the default Kubeconfig
// and writes the result
func WriteMergedDefaultConfig(cfg Config) error {
	pathOptions := clientcmd.NewDefaultPathOptions()
	kubeconfig, err := pathOptions.GetStartingConfig()
	if err != nil {
		return err
	}

	addContainershipConfigAndSetContext(kubeconfig, cfg)

	return clientcmd.ModifyConfig(pathOptions, *kubeconfig, false)
}

// WriteToFile writes a kubeconfig to the given file
func WriteToFile(cfg Config, filename string) error {
	kubeconfig := clientcmdapi.Config{
		Clusters:  map[string]*clientcmdapi.Cluster{},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{},
		Contexts:  map[string]*clientcmdapi.Context{},
	}
	addContainershipConfigAndSetContext(&kubeconfig, cfg)
	return clientcmd.WriteToFile(kubeconfig, filename)
}

// Add a new Kubeconfig to the given starting config for the given Containership config.
// Overwrites any existing Kubeconfig info for the given cluster.
func addContainershipConfigAndSetContext(kubeconfig *clientcmdapi.Config, csConfig Config) {
	clusterName := clusterNameFromID(csConfig.ClusterID)
	userName := userNameFromID(csConfig.UserID)
	contextName := contextNameFromClusterID(csConfig.ClusterID)

	kubeconfig.Clusters[clusterName] = &clientcmdapi.Cluster{
		Server: csConfig.ServerAddress,
	}

	kubeconfig.AuthInfos[userName] = &clientcmdapi.AuthInfo{
		Token: csConfig.Token,
	}

	kubeconfig.Contexts[contextName] = &clientcmdapi.Context{
		Cluster:  clusterName,
		AuthInfo: userName,
	}

	kubeconfig.CurrentContext = contextName
}

func clusterNameFromID(clusterID string) string {
	return fmt.Sprintf("cs-%s", clusterID)
}

func userNameFromID(userID string) string {
	return fmt.Sprintf("cs-%s", userID)
}

func contextNameFromClusterID(clusterID string) string {
	return fmt.Sprintf("cs-ctx-%s", clusterID)
}
