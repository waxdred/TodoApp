package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	projectList    *project
	paginatorIndex int
	spin           spinner.Model
	search         textinput.Model
	searchValue    string
	typing         bool
	projectActive  bool
	projet         lipgloss.Style
	todoActive     bool
	todo           lipgloss.Style
	exitPopup      bool
	err            error
}

func (m *Model) Init() tea.Cmd {
	if m.typing {
		return textinput.Blink
	}
	return m.spin.Tick
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right":
			if !m.exitPopup {
				m.todoActive = true
				m.projectActive = false
			}
			return m, nil
		case "left":
			if !m.exitPopup {
				m.todoActive = false
				m.projectActive = true
			}
			return m, nil
		case "ctrl+r":
			m.typing = true
			m.search.Reset()
		case "esc":
			if m.typing {
				m.typing = false
			} else if m.exitPopup {
				m.exitPopup = false
			} else {
				m.exitPopup = true
			}
		case "enter":
			if m.exitPopup {
				return m, tea.Quit
			}
			if m.typing {
				m.typing = false
			}
		case "tab", "down":
			if m.projectActive {
				if m.projectList.index > m.projectList.size {
					m.projectList.index = 0
				} else {
					m.projectList.index++
				}
			}
		case "up":
			if m.projectActive {
				if m.projectList.index < 0 {
					m.projectList.index = m.projectList.size
				} else {
					m.projectList.index--
				}
			}
		}
		if m.typing {
			var cmd tea.Cmd

			m.search, cmd = m.search.Update(msg)
			if m.search.Value() != "" {
				m.searchValue = m.search.Value()
			}
			return m, cmd
		}
	default:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) View() string {
	var ret string
	var header string
	var k int
	if m.projectActive {
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTabStyle.Render(m.spin.View()+"Project"),
			tabStyle.Render("Todo Project"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)+(width))))
		header = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		pagn := m.projectList.SizePag(SizeList(row))
		if !m.typing {
			var col string
			col = lipgloss.JoinVertical(lipgloss.Top, row, "\n")
			if m.projectList.index >= SizeList(row) {
				k = SizeList(row)
			} else {
				k = 0
			}
			for i := k; i < SizeList(row); i++ {
				if lipgloss.Height(col) < height-6 {
					if i == m.projectList.index {
						col = lipgloss.JoinVertical(lipgloss.Top, col, Selectlist.Render(m.projectList.list[i]))
					} else {
						col = lipgloss.JoinVertical(lipgloss.Top, col, NoSelectlist.Render(m.projectList.list[k]))
					}
					k++
				}
			}
			pagnStr := fmt.Sprint(strings.Repeat("•", pagn))
			col = lipgloss.JoinVertical(lipgloss.Top, col, ActiveDot.Render(pagnStr))
			ret = fmt.Sprintf("%s", col)
		} else if m.typing {
			col := lipgloss.JoinVertical(lipgloss.Top, row, m.search.View())
			for i := 0; i < int(m.projectList.size); i++ {
				col = lipgloss.JoinVertical(lipgloss.Top, col, Selectlist.Render(m.projectList.list[i]))
			}
			ret = fmt.Sprintf("%s", col)
		} else {
			ret = fmt.Sprintf("%s", row)
		}
		m.spin.Tick()
	}
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
		ok := buttonStyle.Render(" ﬌ Enter")
		cancel := buttonStyle.Render("  Esc")
		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Are you sure you want to exit?")
		buttons := lipgloss.JoinHorizontal(lipgloss.Top, ok, "    ", cancel)
		ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)
		dialog := lipgloss.Place(width, height/2,
			lipgloss.Center, lipgloss.Center,
			dialogExit.Render(ui),
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(subtle),
		)
		return header + dialog
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
