package commands

import "github.com/spf13/cobra"

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize stravastats to connect to Strava API on your behalf",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	commands.AddCommand(authCmd)
}
