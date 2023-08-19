package commands

import (
	"stravastats/internal/api"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := api.Request("athlete/activities")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	commands.AddCommand(testCmd)
}
