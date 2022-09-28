package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	TopLeft  = "╭"
	TopRight = "╮"

	BottomLeft  = "╰"
	BottomRight = "╯"

	Middle = "│"
)

func BorderTop(text string) string {
	top := fmt.Sprint(strings.Repeat(" ", max(0, lipgloss.Width(text)+width)))
	return top
}

func BorderBottom(text string) string {
	bottom := fmt.Sprintf(BottomLeft, strings.Repeat(" ", max(0, width-2)), BottomRight)
	return bottom
}

func BorderMiddle(text string) string {
	middle := fmt.Sprintf(Middle, strings.Repeat(" ", max(0, width-2)), Middle)
	return middle
}
