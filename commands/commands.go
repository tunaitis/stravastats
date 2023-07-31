package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var commands = &cobra.Command{
	Use:   "stravastats",
	Short: "stravastats is a CLI utility to show your personal Strava statistics in the terminal",
}

func initConfig() {
	fmt.Println("init config")

	viper.SetConfigName("stravastats")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.stravastats")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if viper.IsSet("clientId") == false {
		log.Fatal("client id wasn't set")
	}

	if viper.IsSet("clientSecret") == false {
		log.Fatal("client secret wasn't set")
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
