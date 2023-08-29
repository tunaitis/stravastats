package cmd

import (
	"fmt"
	"stravastats/internal/service"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		fmt.Println(stats)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
