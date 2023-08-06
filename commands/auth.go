package commands

import (
	"fmt"
	"stravastats/internal/api"
	"stravastats/internal/authcode"
	"stravastats/internal/config"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize stravastats to connect to Strava API on your behalf",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		code, err := authcode.ReadFromStore()
		if err != nil {
			return err
		}

		reauthorize, err := cmd.Flags().GetBool("reauthorize")
		if err != nil {
			return err
		}

		if code != "" && reauthorize == false {
			fmt.Println("You are already authorized. Use the -r flag to reauthorize.")
			return nil
		}

		authUrl, err := api.GetAuthUrl(cfg.Api.ClientId)
		if err != nil {
			return err
		}

		browser.OpenURL(authUrl)

		code, err = authcode.WaitForAuthorizationCode()
		if err != nil {
			return err
		}

		authcode.SaveToStore(code)

		fmt.Println(code)

		return nil
	},
}

func init() {
	commands.AddCommand(authCmd)
	authCmd.Flags().BoolP("reauthorize", "r", false, "reauthorize the API access")
}
