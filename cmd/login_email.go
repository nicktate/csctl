package cmd

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/containership/csctl/cloud/auth"
	"github.com/containership/csctl/cloud/auth/types"
)

// Flags
var (
	email    string
	password string
)

// Questions
var (
	emailQuestion = &survey.Input{
		Message: "email:",
	}

	passwordQuestion = &survey.Password{
		Message: "password:",
	}
)

func emailFlag(cmd *cobra.Command, persistent bool) {
	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVarP(&email, "email", "e", "", "email to log in with")
}

func passwordFlag(cmd *cobra.Command, persistent bool) {
	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVarP(&password, "password", "p", "", "password to log in with")
}

// loginEmailCmd represents the loginEmail command
var loginEmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Log in with an email and password",

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if email == "" {
			err := survey.AskOne(emailQuestion, &email, nil)

			if err != nil {
				return err
			}

			if email == "" {
				return errors.New("You must specify a valid email to log in")
			}
		}

		if password == "" {
			err := survey.AskOne(passwordQuestion, &password, nil)

			if err != nil {
				return err
			}

			if password == "" {
				return errors.New("You must specify a valid password to log in")
			}
		}

		var resp = &types.AccountToken{}
		var err error

		resp, err = clientset.Auth().Login(auth.AuthMethodEmail).Post(&types.LoginRequest{
			Email:    email,
			Password: password,
		})

		if err != nil {
			return err
		}

		viper.Set("token", *resp.Token)

		err = viper.WriteConfig()

		if err != nil {
			return err
		}

		log.Print("Successfully logged in!")

		return nil
	},
}

func init() {
	emailFlag(loginEmailCmd, false)
	passwordFlag(loginEmailCmd, false)
	loginCmd.AddCommand(loginEmailCmd)
}
