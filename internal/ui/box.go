package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Add a purple, rectangular border
var style = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1).
	BorderStyle(lipgloss.RoundedBorder())
	// BorderForeground(lipgloss.Color("63"))

func icon(name string) string {
	switch name {
	case "Run":
		return "🏃"
	case "Ride":
		return "🚴"
	case "Swim":
		return "🏊"
	case "Walk":
		return "🚶"
	}

	return ""
}

func longestLine(lines []string) int {
	l := 0

	for i := range lines {
		if len(lines[i]) > l {
			l = len(lines[i])
		}
	}

	return l
}

func bodyLine(line string, width int) string {
	return ""
}

func Box(name string, distance float32, duration float32) string {
	if distance == 0 {
		return ""
	}

	header := []string{
		fmt.Sprintf("%s %s", icon(name), name),
		"",
	}

	body := []string{
		fmt.Sprintf("Distance %.2f km", distance/1000),
		fmt.Sprintf("Time %.2f h", duration/60/60),
	}

	bodyWidth := longestLine(body)

	body = []string{
		fmt.Sprintf("Distance %*.*f km", (bodyWidth - len("Distance ") - len(" km") + 2), 2, distance/1000),
		fmt.Sprintf("Time %*.*f h", (bodyWidth - len("Time ") - len(" h") + 2), 2, duration/60/60),
	}

	content := strings.Join(header, "\n") + "\n" + strings.Join(body, "\n")

	return style.Render(content)
}
