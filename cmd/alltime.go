package cmd

import (
	"stravastats/internal/config"
	"stravastats/internal/service"
	"stravastats/internal/view"

	"github.com/spf13/cobra"
)

var allTimeCmd = &cobra.Command{
	Use:   "all-time",
	Short: "Show all-time stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		activityFilter := stats.ActivityTypes
		if len(cfg.Display.Activities) > 0 {
			activityFilter = cfg.Display.Activities
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
}
