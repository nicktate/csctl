package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/containership/csctl/cloud"
)

// Flags / config options
var (
	cfgFile      string
	debugEnabled bool

	// TODO support names, not just IDs, and resolve appropriately
	organizationID string
	clusterID      string
	nodePoolID     string

	userToken string
)

var (
	clientset cloud.Interface
)

func orgScopedPreRunE(cmd *cobra.Command, _ []string) error {
	if organizationID != "" {
		// Command line flag was set and it takes precedence
		return nil
	}

	organizationID = viper.GetString("organization")
	if organizationID == "" {
		return errors.New("please specify an organization via --organization or config file")
	}

	return nil
}

func clusterScopedPreRunE(cmd *cobra.Command, _ []string) error {
	if err := orgScopedPreRunE(cmd, nil); err != nil {
		return err
	}

	if clusterID != "" {
		// Command line flag was set and it takes precedence
		return nil
	}

	clusterID = viper.GetString("cluster")
	if clusterID == "" {
		return errors.New("please specify a cluster via --cluster or config file")
	}

	return nil
}

func nodePoolScopedPreRunE(cmd *cobra.Command, _ []string) error {
	if err := clusterScopedPreRunE(cmd, nil); err != nil {
		return err
	}

	if nodePoolID != "" {
		// Command line flag was set and it takes precedence
		return nil
	}

	nodePoolID = viper.GetString("node-pool")
	if nodePoolID == "" {
		return errors.New("please specify a node pool via --node-pool or config file")
	}

	return nil
}

func clientsetRequiredPreRunE(cmd *cobra.Command, _ []string) error {
	if userToken != "" {
		// Command line flag was set and it takes precedence
		return nil
	}

	userToken = viper.GetString("token")
	if userToken == "" {
		return errors.New("please specify a token via --token or config file")
	}

	var err error
	clientset, err = cloud.New(cloud.Config{
		Token:            userToken,
		APIBaseURL:       viper.GetString("apiBaseURL"),
		AuthBaseURL:      viper.GetString("authBaseURL"),
		ProvisionBaseURL: viper.GetString("provisionBaseURL"),
		DebugEnabled:     debugEnabled,
	})

	return err
}

// Note that for all of the flags that we allow to override the config file,
// we cannot use viper.BindPFlag() as that only works for a single flagset
// and we use the same flag across multiple flagsets.
func bindCommandToOrganizationScope(cmd *cobra.Command, persistent bool) {
	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&organizationID, "organization", "", "organization to use")
}

func bindCommandToClusterScope(cmd *cobra.Command, persistent bool) {
	bindCommandToOrganizationScope(cmd, persistent)

	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&clusterID, "cluster", "", "cluster to use")
}

func bindCommandToNodePoolScope(cmd *cobra.Command, persistent bool) {
	bindCommandToClusterScope(cmd, persistent)

	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&nodePoolID, "node-pool", "", "node pool to use")
}

// assumes that if a clientset is required for a command,
// it should be persistent (required for all subcommands)
func requireClientset(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&userToken, "token", "", "Containership Cloud token to use")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "csctl",
	Long: `csctl is a command line interface for Containership.

Find more information at: https://github.com/containership/csctl`,

	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.containership/csctl.yaml)")

	rootCmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, "enable/disable debug mode (trace all HTTP requests)")

	rootCmd.PersistentFlags().StringVar(&userToken, "token", "", "Containership token to authenticate with")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		containershipDir := path.Join(home, ".containership")

		// Search config in ~/.containership directory with name "csctl.yaml"
		viper.AddConfigPath(containershipDir)
		// Note that function expects the extension to be omitted
		viper.SetConfigName("csctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
