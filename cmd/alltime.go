package cmd

import (
	"errors"
	"fmt"
	"stravastats/internal/config"
	"stravastats/internal/service"
	"stravastats/internal/view"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var allTimeCmd = &cobra.Command{
	Use:   "all-time",
	Short: "Show all-time stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			if errors.As(err, &viper.ConfigFileNotFoundError{}) {
				fmt.Println("Config file not found, run `stravastats init` first.")
				return nil
			}

			return err
		}

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		activityFilter, err := cmd.Flags().GetStringSlice("activity")
		if err != nil {
			return err
		}

		if activityFilter == nil || len(activityFilter) == 0 {
			activityFilter = cfg.Display.Activities
		}

		if activityFilter == nil || len(activityFilter) == 0 {
			activityFilter = stats.ActivityTypes
		}

		err = view.AllTime(stats, activityFilter)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(allTimeCmd)
	allTimeCmd.Flags().StringSliceP("activity", "a", nil, "filter activities")
}
