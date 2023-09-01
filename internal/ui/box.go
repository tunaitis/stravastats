package ui

import (
	"fmt"
	"stravastats/internal/model"
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

func Box(activity model.ActivityStats) string {
	if activity.Distance == 0 {
		return ""
	}

	header := []string{
		fmt.Sprintf("%s %s", icon(activity.Type), activity.Type),
		"",
	}

	body := []string{
		fmt.Sprintf("Distance %.2f km", activity.Distance/1000),
		fmt.Sprintf("Time %.2f h", activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, fmt.Sprintf("Elev Gain %.2f h", activity.ElevationGain))
	}

	bodyWidth := longestLine(body)

	body = []string{
		fmt.Sprintf("Distance %*.*f km", (bodyWidth - len("Distance ") - len(" km") + 2), 2, activity.Distance/1000),
		fmt.Sprintf("Time %*.*f h", (bodyWidth - len("Time ") - len(" h") + 2), 2, activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, fmt.Sprintf("Elev Gain %*.*f m", (bodyWidth-len("Elev Gain ")-len(" m")+2), 2, activity.ElevationGain))
	} else {
		body = append(body, "")
	}

	content := strings.Join(header, "\n") + "\n" + strings.Join(body, "\n")

	return style.Render(content)
}
