package ui

import (
	"fmt"

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

func Box(name string, distance float32, duration float32) string {

	if distance == 0 {
		return ""
	}

	distanceLine := fmt.Sprintf("%.2f km", distance/1000)
	durationLine := fmt.Sprintf("%.2f h", duration/60/60)

	content := fmt.Sprintf("%s %s\n\n%s\n%s", icon(name), name, distanceLine, durationLine)

	return style.Render(content)
}
