package cmd

import (
	"stravastats/internal/config"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Open the configuration file with the default text editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath, err := config.GetConfigPath()
		if err != nil {
			return err
		}

		err = open.Run(cfgPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
