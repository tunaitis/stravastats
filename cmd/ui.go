package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "A playground to test various ideas regarding the StravaStats UI",
	RunE: func(cmd *cobra.Command, args []string) error {

		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			PaddingTop(1).
			PaddingBottom(1).
			PaddingLeft(4).
			MarginTop(1).
			MarginRight(1).
			Align(lipgloss.Center).
			Width(22)

		str := lipgloss.JoinHorizontal(0, style.Render("a"), style.Render("b"), style.Render("c"))

		fmt.Println(str)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
