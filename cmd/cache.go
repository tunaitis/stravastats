package cmd

import (
	"stravastats/internal/cache"

	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage stravastats cache",
}

var clearCacheCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the cache",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := cache.RemoveActivities()

		return err
	},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
	cacheCmd.AddCommand(clearCacheCmd)
}
