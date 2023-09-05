package cmd

import (
	"errors"
	"fmt"
	"stravastats/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}

		overwrite, err := cmd.Flags().GetBool("overwrite")
		if err != nil {
			return err
		}

		if (cfg.Api.ClientId != "" || cfg.Api.ClientSecret != "") && overwrite == false {
			return errors.New("client id and secret are already set")
		}

		fmt.Print("Enter Client Id: ")
		fmt.Scanln(&cfg.Api.ClientId)

		fmt.Print("Enter Client Secret: ")
		fmt.Scanln(&cfg.Api.ClientSecret)

		fmt.Println()

		if cfg.Api.ClientId == "" {
			return errors.New("client id can't be empty")
		}

		if cfg.Api.ClientSecret == "" {
			return errors.New("client secret can't be empty")
		}

		err = config.SaveConfig(cfg)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("overwrite", "r", false, "overwrite current configuration file")
}
