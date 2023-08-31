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

func Box(name string, distance float32, duration float32) string {
	if distance == 0 {
		return ""
	}

	header := []string{
		fmt.Sprintf("%s %s", icon(name), name),
		"",
	}

	body := []string{
		fmt.Sprintf("%.2f km", distance/1000),
		fmt.Sprintf("%.2f hh", duration/60/60),
	}

	lw := longestLine(body) + 1

	alignRight := lipgloss.NewStyle().Width(lw).Align(lipgloss.Right)

	for i := range body {
		body[i] = alignRight.Render(body[i])
	}

	content := strings.Join(header, "\n") + "\n" + strings.Join(body, "\n")

	content = strings.ReplaceAll(content, "hh", "h ")

	return style.Render(content)
}
