package cmd

import (
	"fmt"
	"os"
	"stravastats/internal/service"
	"stravastats/internal/ui"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "An empty command for testing",
	RunE: func(cmd *cobra.Command, args []string) error {

		activities := []string{"ride", "run", "swim", "walk"}

		stats, err := service.GetActivityStats()
		if err != nil {
			return err
		}

		var ri = 0
		var rw = 0
		var rows [][]string = [][]string{[]string{}}

		tw, _, err := terminal.GetSize(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}

		for _, k := range activities {
			if a, ok := stats.Activities[k]; ok {
				b := ui.Box(a)
				rw = rw + lipgloss.Width(b)
				if rw > tw {
					rw = 0
					rows = append(rows, []string{})
					ri++
				}

				rows[ri] = append(rows[ri], b)
			}
		}

		//x := lipgloss.JoinHorizontal(lipgloss.Bottom, line...)
		for _, r := range rows {
			x := lipgloss.JoinHorizontal(lipgloss.Bottom, r...)
			fmt.Println(x)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
