package commands

import (
	"fmt"
	"stravastats/internal/api"
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

		authUrl, err := api.GetAuthUrl(cfg.Api.ClientId)
		if err != nil {
			return err
		}

		browser.OpenURL(authUrl)

		code, err := api.WaitForAuthorizationCode()
		if err != nil {
			return err
		}

		fmt.Println(code)

		return nil
	},
}

func init() {
	commands.AddCommand(authCmd)
}
