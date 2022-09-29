package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *Model) Init() tea.Cmd {
	if m.typing || m.projectAdd || m.projectRename {
		return textinput.Blink
	}
	return m.spin.Tick
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
		}
	default:
		return Defatul(m, msg)
	}
	return m, nil
}

func (m *Model) View() string {
	var ret string
	var header string
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
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			tabStyle.Render("Project"),
			activeTabStyle.Render(m.spin.View()+"Todo"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)+(width))))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		ret = fmt.Sprintf("%s", row)
	}
	if m.exitPopup {
		return header + Popup("Are you sure you want to exit?")
	} else if m.DeletePopup {
		return header + Popup("Are you sure you want to Delete?")
	}
	// helper
	helper := fmt.Sprint(
		"\n\n\nAdd: <ctrl+a>   Rename: <ctrl+r>   Delete <ctrl+d>   Search: <ctrl+f>  nav: arrow  exit: <Esc>",
	)
	return ret + helpStyle.Render(helper)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
