package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var commands = &cobra.Command{
	Use:   "stravastats",
	Short: "stravastats is a CLI utility to show your personal Strava statistics in the terminal",
}

func Execute() {
	if err := commands.Execute(); err != nil {
		//fmt.Println(err)
		os.Exit(1)
	}
}
