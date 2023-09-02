package view

import (
	"fmt"
	"os"
	"stravastats/internal/model"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/crypto/ssh/terminal"
)

var style = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingRight(1).
	BorderStyle(lipgloss.RoundedBorder())

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

func cardProp[T int | float32](name string, unit string, width int, value T) string {
	gap := 3

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

func card(activity model.ActivityStats) string {
	if activity.Distance == 0 {
		return ""
	}

	header := []string{
		fmt.Sprintf("%s %s", icon(activity.Type), activity.Type),
		"",
	}

	body := []string{
		cardProp("Activities", "", 0, activity.Count),
		cardProp("Distance", "km", 0, activity.Distance/1000),
		cardProp("Time", "h", 0, activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, cardProp("Elev Gain", "m", 0, activity.ElevationGain))
	}

	bodyWidth := longestLine(body)

	body = []string{
		cardProp("Activities", "", bodyWidth, activity.Count),
		cardProp("Distance", "km", bodyWidth, activity.Distance/1000),
		cardProp("Time", "h", bodyWidth, activity.Duration/60/60),
	}

	if activity.Type != "Swim" {
		body = append(body, cardProp("Elev Gain", "m", bodyWidth, activity.ElevationGain))
	} else {
		body = append(body, "")
	}

	content := strings.Join(header, "\n") + "\n" + strings.Join(body, "\n")

	return style.Render(content)
}

func AllTime(stats *model.Stats) error {
	activities := []string{"ride", "run", "swim", "walk"}

	var ri = 0
	var rw = 0
	var rows [][]string = [][]string{[]string{}}

	tw, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	for _, k := range activities {
		if a, ok := stats.Activities[k]; ok {
			b := card(a)
			rw = rw + lipgloss.Width(b)
			if rw > tw {
				rw = 0
				rows = append(rows, []string{})
				ri++
			}

			rows[ri] = append(rows[ri], b)
		}
	}

	for _, r := range rows {
		x := lipgloss.JoinHorizontal(lipgloss.Bottom, r...)
		fmt.Println(x)
	}

	return nil
}
