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

var gap int = 3

func field[T int | float32](name string, unit string, width int, value T) string {
	nameWidth := len(name) + 1
	unitWidth := 0
	if len(unit) > 0 {
		unitWidth = len(unit) + 1
		unit = " " + unit
	}

	if width == 0 {
		switch any(value).(type) {
			case int:
			return fmt.Sprintf("%s %d%s", name, value, unit)
			case float32:
			return fmt.Sprintf("%s %.2f%s", name, value, unit)
		}
	}

	switch any(value).(type) {
		case int:
		return fmt.Sprintf("%s %*d%s", name, (width - nameWidth - unitWidth + gap), value, unit)
		case float32:
		return fmt.Sprintf("%s %*.2f%s", name, (width - nameWidth - unitWidth + gap), value, unit)
	}

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
		field("Activities", "", 0, activity.Count),
		field("Distance", "km", 0, activity.Distance/1000),
		field("Time", "h", 0, activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, field("Elev Gain", "m", 0, activity.ElevationGain))
	}

	bodyWidth := longestLine(body)

	body = []string{
		field("Activities", "", bodyWidth, activity.Count),
		field("Distance", "km", bodyWidth, activity.Distance/1000),
		field("Time", "h", bodyWidth, activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, field("Elev Gain", "m", bodyWidth, activity.ElevationGain))
	} else {
		body = append(body, "")
	}

	content := strings.Join(header, "\n") + "\n" + strings.Join(body, "\n")

	return style.Render(content)
}
