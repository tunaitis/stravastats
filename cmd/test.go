package cmd

import (
	"fmt"
	"stravastats/internal/service"
	"stravastats/internal/ui"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		activities := []string{"ride", "run", "swim"}

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		var line []string

		for _, k := range activities {
			if a, ok := stats.Activities[k]; ok {
				b := ui.Box(a.Type, a.Distance)
				line = append(line, b)
			}
		}

		x := lipgloss.JoinHorizontal(lipgloss.Bottom, line...)

		fmt.Println(x)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
