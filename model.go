package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) Init() tea.Cmd {
	// if m.typing || m.projectAdd || m.projectRename {
	// 	return textinput.Blink
	// } else if m.textareaActive {
	// 	return textarea.Blink
	// }
	return tea.Batch(textarea.Blink, textinput.Blink, spinner.Tick)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			return right(m)
		case "left":
			return left(m)
		case "ctrl+f":
			m.typing = true
		case "esc":
			esc(m)
		case "enter":
			return enter(m)
		case "tab", "down", "j":
			down(m)
		case "up", "k":
			up(m)
		case "ctrl+a":
			ctrla(m)
		case "ctrl+r":
			ctrlr(m)
		case "ctrl+d":
			ctrld(m)
		}
		if m.typing {
			return ActiveProjectSelect(m, msg)
		} else if m.projectAdd {
			return ActiveProjectAdd(m, msg)
		} else if m.projectRename {
			return ActiveProjectRename(m, msg)
		} else if m.textareaActive {
			m.textarea, _ = m.textarea.Update(msg)
		}
	default:
		if !m.textarea.Focused() {
			cmd := m.textarea.Focus()
			return m, cmd
		}
		return Defatul(m, msg)
	}
	return m, nil
}

func (m *Model) View() string {
	var ret string
	var header string
	var helper string
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		activeTabStyle.Render(m.spin.View()+"Project"),
		tabStyle.Render("Todo Project"),
	)
	if m.projectActive {
		ret = ProjectActive(m, row, header, ret)
	} else {
		ret = fmt.Sprintf("%s", row)
	}
	m.spin.Tick()
	if m.todoActive {
		ret = todoView(m, ret)
	}
	if m.exitPopup {
		return header + Popup("Are you sure you want to exit?")
	} else if m.DeletePopup {
		return header + Popup("Are you sure you want to Delete?")
	} else if m.AddPopup {
		return header + PopupAdd("Todo exist already!")
	}
	// helper
	if m.projectActive {
		helper = fmt.Sprint(
			"\n\n\nAdd: <ctrl+a>   Rename: <ctrl+r>   Delete <ctrl+d>   Search: <ctrl+f>  nav: arrow  exit: <Esc>",
		)
	}
	return ret + helpStyle.Render(helper)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
