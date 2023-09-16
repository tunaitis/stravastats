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
			err := config.SetValue(&cfg, args[0], args[1])
			if err != nil {
				return err
			}
			err = config.SaveConfig(cfg)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("edit", "e", false, "open the configuration file with the default editor")
}
