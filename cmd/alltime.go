package cmd

import (
	"stravastats/internal/service"
	"stravastats/internal/view"

	"github.com/spf13/cobra"
)

var allTimeCmd = &cobra.Command{
	Use:   "all-time",
	Short: "Show all-time stats",
	RunE: func(cmd *cobra.Command, args []string) error {

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		err = view.AllTime(stats)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(allTimeCmd)
}
