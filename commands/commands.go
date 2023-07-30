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

	if viper.IsSet("client_id") == false {
		log.Fatal("client_id wasn't set")
	}

	fmt.Printf("client id: %s\n", viper.GetString("client_id"))
	fmt.Printf("client secred: %s\n\n", viper.GetString("client_secret"))

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
