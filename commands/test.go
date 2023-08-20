package commands

import (
	"fmt"
	"stravastats/internal/api"
	"stravastats/internal/models"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		activities, err := api.Request[[]models.Activity]("athlete/activities", nil)
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
