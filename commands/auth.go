package commands

import (
	"fmt"
	"stravastats/internal/api"
	"stravastats/internal/config"
	"stravastats/internal/keychain"

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

		accessToken, refreshToken, err := keychain.ReadTokens()
		if err != nil {
			return err
		}

		reauthorize, err := cmd.Flags().GetBool("reauthorize")
		if err != nil {
			return err
		}

		if accessToken != "" && refreshToken != "" && reauthorize == false {
			fmt.Println("You are already authorized. Use the -r flag to reauthorize.")
			return nil
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

		accessToken, refreshToken, err = api.ExchangeCodeToAccessToken(cfg.Api.ClientId, cfg.Api.ClientSecret, code)
		if err != nil {
			return err
		}

		err = keychain.WriteTokens(accessToken, refreshToken)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	commands.AddCommand(authCmd)
	authCmd.Flags().BoolP("reauthorize", "r", false, "reauthorize the API access")
}
