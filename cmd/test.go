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

		types, err := service.GetActivityTypes()
		if err != nil {
			return err
		}

		fmt.Println(types)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
