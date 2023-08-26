package cmd

import (
	"stravastats/internal/api"
	"stravastats/internal/cache"
	"time"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		cached := cache.GetActivities()

		return nil
		from := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		if len(cached) > 0 {
			from = cached[0].StartDate
		}

		activities, err := api.GetActivities(from)
		if err != nil {
			return err
		}

		if len(activities) > 0 {
			cached = append(cached, activities...)
		}

		err = cache.SetActivities(cached)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	commands.AddCommand(testCmd)
}
