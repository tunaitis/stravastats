package cmd

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

		tokens, err := config.ReadTokens()
		if err != nil {
			return err
		}

		deleteTokens, err := cmd.Flags().GetBool("delete")
		if err != nil {
			return err
		}

		if deleteTokens {
			err = config.DeleteTokens()
			if err != nil {
				return err
			}

			fmt.Println("Tokens have been deleted.")

			return nil
		}

		reauthorize, err := cmd.Flags().GetBool("reauthorize")
		if err != nil {
			return err
		}

		if tokens.AccessToken != "" && tokens.RefreshToken != "" && reauthorize == false {
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

		tokens, err = api.ExchangeCodeToAccessToken(cfg.Api.ClientId, cfg.Api.ClientSecret, code)
		if err != nil {
			return err
		}

		err = config.SaveTokens(tokens)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().BoolP("reauthorize", "r", false, "reauthorize the API access")
	authCmd.Flags().BoolP("delete", "d", false, "delete saved tokens from the keyring")
}
