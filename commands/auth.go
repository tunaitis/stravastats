package commands

import (
	"fmt"
	"stravalog/internal/config"

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

		var authUrl = fmt.Sprintf("https://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost/exchange_token&approval_prompt=force&scope=read", cfg.Api.ClientId)
		browser.OpenURL(authUrl)

		return nil
	},
}

func init() {
	commands.AddCommand(authCmd)
}
