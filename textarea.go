package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/lipgloss"
)

func InittextArea() textarea.Model {
	ta := textarea.New()

	ta.Placeholder = "Description Todo"
	ta.Focus()
	ta.FocusedStyle = textarea.Style{
		Text:       TextStyle,
		Prompt:     CursorStyle,
		CursorLine: TextStyle,
	}
	ta.Prompt = "┃ "
	ta.CharLimit = 280
	ta.SetWidth(width / 3)
	ta.SetHeight(5)
	ta.ShowLineNumbers = true
	return ta
}

func ViewTextAre(m *Model, name string) string {
	title := TextStyleInput.Render("Title: ")
	desc := TextStyleInput.Render("Description: ")

	ui := lipgloss.JoinVertical(lipgloss.Top, title, m.search.View(), desc, m.textarea.View())
	box := lipgloss.Place(
		width,
		height/2,
		lipgloss.Center,
		lipgloss.Center,
		dialogExit.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
	)
	return box
}
