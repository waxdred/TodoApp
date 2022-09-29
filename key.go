package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func right(m *Model) (tea.Model, tea.Cmd) {
	if !m.exitPopup {
		m.todoActive = true
		m.projectActive = false
	}
	return m, nil
}

func left(m *Model) (tea.Model, tea.Cmd) {
	if !m.exitPopup {
		m.todoActive = false
		m.projectActive = true
	}
	return m, nil
}

func esc(m *Model) (tea.Model, tea.Cmd) {
	if m.typing {
		m.typing = false
	} else if m.exitPopup {
		m.exitPopup = false
	} else if m.DeletePopup {
		m.DeletePopup = false
	} else {
		m.exitPopup = true
	}
	return m, nil
}

func enter(m *Model) (tea.Model, tea.Cmd) {
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
	return m, nil
}

func down(m *Model) (tea.Model, tea.Cmd) {
	if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
		if m.projectList.index > m.projectList.size {
			m.projectList.index = 0
		} else {
			m.projectList.index++
		}
	}
	return m, nil
}

func up(m *Model) (tea.Model, tea.Cmd) {
	if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
		if m.projectList.index < 0 {
			m.projectList.index = m.projectList.size
		} else {
			m.projectList.index--
		}
	}
	return m, nil
}

func ctrla(m *Model) (tea.Model, tea.Cmd) {
	if !m.projectAdd {
		m.addbuffer.Blink()
		m.projectAdd = true
	}
	return m, nil
}

func ctrlr(m *Model) (tea.Model, tea.Cmd) {
	if !m.projectRename {
		m.renamebuffer.Blink()
		m.projectRename = true
	}
	return m, nil
}

func ctrld(m *Model) (tea.Model, tea.Cmd) {
	if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
		m.DeletePopup = true
	}
	return m, nil
}

func ActiveProjectSelect(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.search, cmd = m.search.Update(msg)
	if m.search.Value() != "" {
		m.searchValue = m.search.Value()
	}
	return m, cmd
}

func ActiveProjectAdd(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.addbuffer, cmd = m.addbuffer.Update(msg)
	if m.addbuffer.Value() != "" {
		m.addValue = m.addbuffer.Value()
	}
	return m, cmd
}

func ActiveProjectRename(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.renamebuffer, cmd = m.renamebuffer.Update(msg)
	if m.renamebuffer.Value() != "" {
		m.RenameValue = m.renamebuffer.Value()
	}
	return m, cmd
}

func Defatul(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spin, cmd = m.spin.Update(msg)
	return m, cmd
}

func ProjectActive(m *Model, row, header, ret string) string {
	var k int
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
	return ret
}
