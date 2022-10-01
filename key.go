package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func right(m *Model) (tea.Model, tea.Cmd) {
	if m.todoActive {
		m.todoView++
		if m.todoView > 2 {
			m.todoView = 0
		}
	}
	return m, nil
}

func left(m *Model) (tea.Model, tea.Cmd) {
	if m.todoActive {
		m.todoView--
		if m.todoView < 0 {
			m.todoView = 2
		}
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
	} else if m.AddPopup {
		m.AddPopup = false
	} else if m.projectRename {
		m.projectRename = false
	} else if m.todoActive {
		m.todoActive = false
		m.projectActive = true
		m.projectList.index = 0
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
		ret := m.projectList.Add(m.addValue)
		if !ret {
			m.AddPopup = true
		}
		m.addbuffer.Reset()
		return m, nil
	}
	if m.AddPopup {
		m.AddPopup = false
		return m, nil
	}
	if m.projectRename {
		m.projectRename = false
		m.projectList.Rename(m.RenameValue, m.projectList.index)
		m.renamebuffer.Reset()
		return m, nil
	}
	if m.DeletePopup {
		m.projectList.Delete(m.projectList.index)
		m.DeletePopup = false
		return m, nil
	}
	if m.projectActive {
		m.todoActive = true
		m.projectActive = false
		m.Todo.GetTodo(m.projectList.list[m.projectList.index])
		return m, nil
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
	} else if m.todoActive && m.todoView == 0 && m.Todo.Todo.Len > 1 {
		if m.Todo.Todo.Idx >= m.Todo.Todo.Len-1 {
			m.Todo.Todo.Idx = 0
		} else {
			m.Todo.Todo.Idx++
		}
	} else if m.todoActive && m.todoView == 1 && m.Todo.Progress.Len > 1 {
		if m.Todo.Progress.Idx >= m.Todo.Progress.Len-1 {
			m.Todo.Progress.Idx = 0
		} else {
			m.Todo.Progress.Idx++
		}
	} else if m.todoActive && m.todoView == 2 && m.Todo.Finish.Len > 1 {
		if m.Todo.Finish.Idx >= m.Todo.Finish.Len-1 {
			m.Todo.Finish.Idx = 0
		} else {
			m.Todo.Finish.Idx++
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
	} else if m.todoActive && m.todoView == 0 && m.Todo.Todo.Len > 1 {
		if m.Todo.Todo.Idx == 0 {
			m.Todo.Todo.Idx = m.Todo.Todo.Len - 1
		} else {
			m.Todo.Todo.Idx--
		}
	} else if m.todoActive && m.todoView == 1 && m.Todo.Progress.Len > 1 {
		if m.Todo.Progress.Idx == 0 {
			m.Todo.Progress.Idx = m.Todo.Progress.Len - 1
		} else {
			m.Todo.Progress.Idx--
		}
	} else if m.todoActive && m.todoView == 2 && m.Todo.Finish.Len > 1 {
		if m.Todo.Finish.Idx == 0 {
			m.Todo.Finish.Idx = m.Todo.Finish.Len - 1
		} else {
			m.Todo.Finish.Idx--
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
