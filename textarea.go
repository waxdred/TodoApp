package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/lipgloss"
)

func InittextArea() PopTodo {
	ta := textarea.New()
	ti := InputInit(tabTodoNoSelect, "Tiltle")

	ta.Placeholder = "Description Todo"
	ta.Blur()
	ti.Blur()
	ta.FocusedStyle = textarea.Style{
		Text:       TextStyle,
		Prompt:     CursorStyle,
		CursorLine: TextStyle,
	}
	ta.Prompt = "┃ "
	ta.CharLimit = 280
	ta.SetWidth(width / 5)
	ta.SetHeight(5)
	ta.ShowLineNumbers = true
	return PopTodo{
		textarea:       ta,
		input:          ti,
		textareaActive: false,
		inputActive:    false,
		confirm:        0,
	}
}

func ViewTextAre(m *Model, name string) string {
	var entry string
	var cancel string
	if m.PopTodo.confirm == 2 {
		entry = buttonStyleSelect.Render(" ﬌ Enter")
	} else {
		entry = buttonStyle.Render(" ﬌ Enter")
	}
	if m.PopTodo.confirm == 3 {
		cancel = buttonStyleSelect.Render("   Cancel")
	} else {
		cancel = buttonStyle.Render("   Cancel")
	}
	title := TextStyleInput.Render("Title: ")
	desc := TextStyleInput.Render("Description: ")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, entry, "    ", cancel)

	ui := lipgloss.JoinVertical(lipgloss.Top, title, m.PopTodo.input.View(), " ", desc, m.PopTodo.textarea.View(), buttons)
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
