package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func InputInit(style lipgloss.Style) textinput.Model {
	t := textinput.NewModel()
	t.CharLimit = 30
	t.CursorStyle = cursorStyle
	t.TextStyle = style
	t.Placeholder = " Search "
	t.Focus()
	return t
}

func SpinInit(style lipgloss.Style) spinner.Model {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = style
	return s
}
