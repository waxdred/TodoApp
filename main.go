package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	spinner   spinner.Model
	textInput textinput.Model
	tying     bool
	loading   bool
	err       error
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.tying = false
			m.loading = true
			return m, spinner.Tick
		}
	}
	if m.tying {
		var cmd tea.Cmd
		var cmdSpin tea.Cmd

		m.spinner, cmdSpin = m.spinner.Update(msg)

		m.textInput, cmd = m.textInput.Update(msg)
		return m, tea.Batch(cmdSpin, cmd)
	}
	if m.loading {
		var cmd tea.Cmd

		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) View() string {
	if m.tying {
		return fmt.Sprintf("%s Enter:\n%s", m.spinner.View(), m.textInput.View())
	}
	if m.loading {
		return fmt.Sprintf("%s ...", m.spinner.View())
	}
	return "ctrl + c for exit"
}

func main() {
	t := textinput.NewModel()
	t.Focus()

	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	initModel := Model{
		spinner:   s,
		textInput: t,
		tying:     true,
	}
	errTea := tea.NewProgram(&initModel, tea.WithAltScreen()).Start()
	if errTea != nil {
		fmt.Println("Error: ", errTea)
		os.Exit(1)
	}
}
