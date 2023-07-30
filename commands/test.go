package commands

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	commands.AddCommand(testCmd)
}
