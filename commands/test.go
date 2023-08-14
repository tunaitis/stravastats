package commands

import (
	"stravastats/internal/api"
	"stravastats/internal/config"
	"stravastats/internal/keychain"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		accessToken, refreshToken, err := keychain.ReadTokens()
		if err != nil {
			return err
		}

		accessToken, refreshToken, err = api.RefreshAccessToken(cfg.Api.ClientId, cfg.Api.ClientSecret, refreshToken)

		client := api.NewClient(accessToken, refreshToken)
		client.ListActivities()

		return nil
	},
}

func init() {
	commands.AddCommand(testCmd)
}
