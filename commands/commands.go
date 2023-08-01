package commands

import (
	"fmt"
	"log"
	"os"
	"stravalog/internal/config"

	"github.com/spf13/cobra"
)

var commands = &cobra.Command{
	Use:   "stravastats",
	Short: "stravastats is a CLI utility to show your personal Strava statistics in the terminal",
}

type Config struct {
	Api ApiConfig `mapstructure:"api"`
}

type ApiConfig struct {
	ClientId     string
	ClientSecret string
}

func initConfig() {
	_, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("ann error has occurred: %s", err)
	}
}

func Execute() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
