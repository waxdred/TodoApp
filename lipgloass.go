package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	widthSearch = 30

	columnWidth = 30
	divisor     = 3
)

var width, height = InitSizeWindow()

// style border
var (
	// Color
	blue      = "#268BD2"
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: blue, Dark: blue}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// helper
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFAB00")).
			Margin(1, 2)

	// cursor
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	cursorStyle  = focusedStyle.Copy()
	// text area
	TextStyleInput = lipgloss.NewStyle().Foreground(subtle).Padding(1, 1)
	TextStyle      = lipgloss.NewStyle().Foreground(subtle)
	CursorStyle    = lipgloss.NewStyle().Foreground(highlight)
	CursorLine     = lipgloss.NewStyle().Foreground(lipgloss.Color("#bcbcbc"))
	// spinner
	spinStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(blue))

	// paginator
	ActiveDot   = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"})
	InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"})

	// todo style
	tabTodoStyle = lipgloss.NewStyle().
			Border(borderRadius).
			Width(width/3).
			Padding(0, 1).Margin(0, 1)
	// select tab
	tabTodoSelect       = tabTodoStyle.Copy().BorderForeground(lipgloss.Color(blue))
	tabTodoSelectTop    = tabTodoSelect.Copy().BorderBottom(false)
	tabTodoSelectMiddle = tabTodoSelect.Copy().BorderBottom(false).BorderTop(false)
	tabTodoSelectBottom = tabTodoSelect.Copy().BorderTop(false)

	// no select tab
	tabTodoNoSelect       = tabTodoStyle.Copy().BorderForeground(subtle)
	tabTodoNoSelectTop    = tabTodoNoSelect.Copy().BorderBottom(false)
	tabTodoNoSelectMiddle = tabTodoNoSelect.Copy().BorderBottom(false).BorderTop(false)
	tabTodoNoSelectBottom = tabTodoNoSelect.Copy().BorderTop(false)

	tabTodoStyleActive = tabTodoStyle.Copy().BorderForeground(lipgloss.Color(blue))
	taskSelect         = lipgloss.NewStyle().Foreground(lipgloss.Color(blue))
	taskNoSelect       = lipgloss.NewStyle().Foreground(subtle)

	// Description todo
	DescriptionStyle = lipgloss.NewStyle().
				Border(borderRadius).
				Width(width+7).
				Padding(0, 1).Margin(0, 1)
	DescriptionSelectTop    = DescriptionStyle.Copy().BorderBottom(false)
	DescriptionSelectMiddle = DescriptionStyle.Copy().BorderBottom(false).BorderTop(false)
	DescriptionSelectBottom = DescriptionStyle.Copy().BorderTop(false)
	DescTiltleStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color(blue)).
				Bold(true).Underline(true).Padding(1, 0)
	DescNorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(blue)).
			Bold(true).Underline(true).Padding(0, 0)
	DescStyle     = lipgloss.NewStyle().Margin(0, 4)
	DescDateStyle = lipgloss.NewStyle().Margin(0, 1)

	// border
	borderActive = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	borderRadius = lipgloss.Border(lipgloss.RoundedBorder())

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tabGap = tabStyle.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// tab style
	tabStyle = lipgloss.NewStyle().
			Border(tabBorder, true).
			Foreground(highlight).
			BorderForeground(lipgloss.Color(blue)).
			Padding(0, 1)

	activeTabStyle = lipgloss.NewStyle().
			Border(borderActive, true).
			Foreground(subtle).
			BorderForeground(lipgloss.Color(blue)).
			Padding(0, 1)

	titleStype = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	searchStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1)

	// exit box
	dialogExit = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(blue)).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			BorderBackground(lipgloss.Color("#888B7E")).
			Padding(0, 3).Margin(1, 10)
	buttonStyleSelect = lipgloss.NewStyle().
				Foreground(lipgloss.Color(blue)).
				Background(lipgloss.Color("#888B7E")).
				BorderBackground(lipgloss.Color("#888B7E")).
				Padding(0, 3).Margin(1, 10)
	buttonStyleAdd = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			BorderBackground(lipgloss.Color("#888B7E")).
			Padding(0, 3).MarginTop(1)

	// Project list
	Selectlist = lipgloss.NewStyle().
			Border(borderRadius).
			BorderForeground(lipgloss.Color(blue)).
			Foreground(highlight).
			Margin(0).
			Padding(0, 1).
			Width(width / 4)

	NoSelectlist = lipgloss.NewStyle().
			Border(borderRadius).
			Foreground(subtle).
			Padding(0, 1).
			Width(width / 4)
)

func InitSizeWindow() (int, int) {
	var width, height int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	stdout, err := cmd.Output()
	if err != nil {
		return 0, 0
	}
	str := string(stdout)
	cmds := strings.Split(str, " ")
	cmds[0] = strings.Replace(cmds[0], "\n", "", -1)
	cmds[1] = strings.Replace(cmds[1], "\n", "", -1)
	width, _ = strconv.Atoi(cmds[1])
	height, _ = strconv.Atoi(cmds[0])
	return width - 20, height - 4
}
