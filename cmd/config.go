package cmd

import (
	"fmt"
	"stravastats/internal/config"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config name value",
	Short: "Get and set configuration properties",
	RunE: func(cmd *cobra.Command, args []string) error {

		edit, err := cmd.Flags().GetBool("edit")
		if err != nil {
			return err
		}

		if edit {
			cfgPath, err := config.GetConfigPath()
			if err != nil {
				return err
			}

			err = open.Run(cfgPath)
			if err != nil {
				return err
			}

			return nil
		}

		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		if len(args) == 1 {
			val, err := config.GetValue(cfg, args[0])
			if err != nil {
				return err
			}

			fmt.Println(val)

			return nil
		}

		if len(args) == 2 {
			changed := false

			if args[0] == "api.clientId" {
				cfg.Api.ClientId = args[1]
				changed = true
			}

			if args[0] == "api.clientSecret" {
				cfg.Api.ClientSecret = args[1]
				changed = true
			}

			if changed {
				err = config.SaveConfig(cfg)
				if err != nil {
					return err
				}

				return nil
			}

			return fmt.Errorf("variable not found: %s", args[0])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("edit", "e", false, "open the configuration file with the default editor")
}
