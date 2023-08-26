package commands

import (
	"os"
	"stravastats/internal/util"

	"github.com/spf13/cobra"
)

var debug bool

var commands = &cobra.Command{
	Use:   "stravastats",
	Short: "stravastats is a CLI utility to show your personal Strava statistics in the terminal",
}

func initConfig() {
	if debug {
		util.ConfigureLogging()
	}
}

func Execute() {
	commands.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Show debug information")

	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
