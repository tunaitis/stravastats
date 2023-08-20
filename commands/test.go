package commands

import (
	"fmt"
	"stravastats/internal/api"
	"time"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		from := time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC)

		activities, err := api.GetActivities(from)
		if err != nil {
			return err
		}

		fmt.Println(activities)

		return nil
	},
}

func init() {
	commands.AddCommand(testCmd)
}
