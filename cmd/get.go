package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// Flags
var (
	outputFormat string
	ownerID      string
	mineOnly     bool
)

func outputResponse(d resource.Displayable, listView bool) {
	var err error
	switch {
	case outputFormat == "" || outputFormat == "table":
		err = d.Table(os.Stdout)

	case outputFormat == "json":
		err = d.JSON(os.Stdout, listView)

	case outputFormat == "yaml":
		err = d.YAML(os.Stdout, listView)

	case strings.HasPrefix(outputFormat, "jsonpath"):
		fields := strings.SplitN(outputFormat, "=", 2)
		if len(fields) != 2 {
			err = errors.New("please specify jsonpath using -ojsonpath=<path>")
			break
		}

		template := fields[1]
		err = d.JSONPath(os.Stdout, template)

	default:
		// TODO handle this using cobra itself?
		err = errors.Errorf("output format %s not supported", outputFormat)
	}

	if err != nil {
		fmt.Println(err)
	}
}

func getMyAccountID() (string, error) {
	me, err := clientset.API().Account().Get()
	if err != nil {
		return "", errors.Wrap(err, "retrieving account")
	}

	return string(me.ID), nil
}

func getPreRunE(cmd *cobra.Command, args []string) error {
	err := clientsetRequiredPreRunE(cmd, args)

	if err != nil {
		return err
	}

	if mineOnly && ownerID != "" {
		return errors.New("you may not specify both --mine and --owner-id simultaneously")
	}

	if mineOnly {
		me, err := getMyAccountID()
		if err != nil {
			return err
		}

		ownerID = me
	}

	return nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource",

	PersistentPreRunE: getPreRunE,
}

func init() {
	rootCmd.AddCommand(getCmd)

	requireClientset(getCmd)

	getCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "output format")
	getCmd.PersistentFlags().StringVar(&ownerID, "owner-id", "", "only list resources owned by the given owner-id")
	getCmd.PersistentFlags().BoolVarP(&mineOnly, "mine", "m", false, "only list resources your user owns")
}
