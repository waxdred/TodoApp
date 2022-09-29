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

func PopupAdd(text string) string {
	ok := buttonStyle.Render(" ﬌ Enter")
	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(text)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, ok)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
	dialog := lipgloss.Place(width, height/2,
		lipgloss.Center, lipgloss.Center,
		dialogExit.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	return dialog
}

func ActivatePagn(pagn, size, idx int, col string) string {
	var row string
	i := 1
	for k := 0; k < pagn; k++ {
		if idx < size*i {
			break
		}
		i++
	}
	for k := 0; k < pagn; k++ {
		if k == i-1 {
			row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, InactiveDot.Render("•"))
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, ActiveDot.Render("•"))
		}
	}
	col = lipgloss.JoinVertical(lipgloss.Top, col, row)
	return col
}
