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
	projectAdd     bool
	projectRename  bool
	projectList    *project
	paginatorIndex int
	spin           spinner.Model
	search         textinput.Model
	addbuffer      textinput.Model
	renamebuffer   textinput.Model
	searchValue    string
	addValue       string
	RenameValue    string
	typing         bool
	projectActive  bool
	projet         lipgloss.Style
	todoActive     bool
	todo           lipgloss.Style
	exitPopup      bool
	DeletePopup    bool
	err            error
}

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
		case "ctrl+f":
			m.typing = true
		case "esc":
			if m.typing {
				m.typing = false
			} else if m.exitPopup {
				m.exitPopup = false
			} else if m.DeletePopup {
				m.DeletePopup = false
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
			if m.projectAdd {
				m.projectAdd = false
				m.projectList.Add(m.addValue)
				m.addbuffer.Reset()
			}
			if m.projectRename {
				m.projectRename = false
				m.projectList.Rename(m.RenameValue, m.projectList.index)
				m.renamebuffer.Reset()
			}
			if m.DeletePopup {
				m.projectList.Delete(m.projectList.index)
				m.DeletePopup = false
			}
		case "tab", "down", "j":
			if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
				if m.projectList.index > m.projectList.size {
					m.projectList.index = 0
				} else {
					m.projectList.index++
				}
			}
		case "up", "k":
			if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
				if m.projectList.index < 0 {
					m.projectList.index = m.projectList.size
				} else {
					m.projectList.index--
				}
			}
		case "ctrl+a":
			if !m.projectAdd {
				m.addbuffer.Blink()
				m.projectAdd = true
			}
		case "ctrl+r":
			if !m.projectRename {
				m.renamebuffer.Blink()
				m.projectRename = true
			}
		case "ctrl+d":
			if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
				m.DeletePopup = true
			}
		}
		if m.typing {
			var cmd tea.Cmd

			m.search, cmd = m.search.Update(msg)
			if m.search.Value() != "" {
				m.searchValue = m.search.Value()
			}
			return m, cmd
		} else if m.projectAdd {
			var cmd tea.Cmd

			m.addbuffer, cmd = m.addbuffer.Update(msg)
			if m.addbuffer.Value() != "" {
				m.addValue = m.addbuffer.Value()
			}
			return m, cmd
		} else if m.projectRename {
			var cmd tea.Cmd

			m.renamebuffer, cmd = m.renamebuffer.Update(msg)
			if m.renamebuffer.Value() != "" {
				m.RenameValue = m.renamebuffer.Value()
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
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		activeTabStyle.Render(m.spin.View()+"Project"),
		tabStyle.Render("Todo Project"),
	)
	if m.projectActive {
		gap := tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)+(width))))
		header = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		pagn := m.projectList.SizePag(SizeList(row))
		var col string
		if m.typing {
			col = lipgloss.JoinVertical(lipgloss.Top, row, m.search.View())
		} else if m.projectAdd {
			col = lipgloss.JoinVertical(lipgloss.Top, row, m.addbuffer.View())
		} else if m.projectRename {
			col = lipgloss.JoinVertical(lipgloss.Top, row, m.renamebuffer.View())
		} else {
			col = lipgloss.JoinVertical(lipgloss.Top, row, "\n")
		}
		if m.projectList.index >= SizeList(row) {
			k = SizeList(row)
		} else {
			k = 0
		}
		for i := 0; i < SizeList(row); i++ {
			if lipgloss.Height(col) < height-6 && k < m.projectList.size {
				if k == m.projectList.index {
					col = lipgloss.JoinVertical(lipgloss.Top, col, Selectlist.Render(m.projectList.list[k]))
				} else if k < m.projectList.size {
					col = lipgloss.JoinVertical(lipgloss.Top, col, NoSelectlist.Render(m.projectList.list[k]))
				}
				k++
			} else if lipgloss.Height(col) < height-6 {
				col = lipgloss.JoinVertical(lipgloss.Top, col, "\n\n")
			}
		}
		pagnStr := fmt.Sprint(strings.Repeat("•", pagn))
		col = lipgloss.JoinVertical(lipgloss.Top, col, ActiveDot.Render(pagnStr))
		ret = fmt.Sprintf("%s", col)
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
