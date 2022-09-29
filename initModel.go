package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func InputInit(style lipgloss.Style, holder string) textinput.Model {
	t := textinput.NewModel()
	t.CharLimit = 30
	t.CursorStyle = cursorStyle
	t.TextStyle = style
	t.Placeholder = holder
	t.Focus()
	return t
}

func SpinInit(style lipgloss.Style) spinner.Model {
	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = style
	return s
}

func Popup(text string) string {
	ok := buttonStyle.Render(" ﬌ Enter")
	cancel := buttonStyle.Render("  Esc")
	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(text)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, ok, "    ", cancel)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
	dialog := lipgloss.Place(width, height/2,
		lipgloss.Center, lipgloss.Center,
		dialogExit.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	return dialog
}
