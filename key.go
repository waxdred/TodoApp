package main

import (
	"github.com/charmbracelet/bubbles/textarea"
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

func ctrlp(m *Model) (tea.Model, tea.Cmd) {
	if m.todoActive && m.todoView != 1 {
		if m.todoView == 0 {
			m.Todo.AddProgress(m.Todo.Todo.Title[m.Todo.Todo.Idx].Title, m.Todo.Todo.Desc[m.Todo.Todo.Idx].Title)
			m.Todo.Delete(m.Todo.Todo.Idx+1, m.todoView)
			m.Todo.Update()
		} else if m.todoView == 2 {
			m.Todo.AddProgress(m.Todo.Finish.Title[m.Todo.Finish.Idx].Title, m.Todo.Finish.Desc[m.Todo.Finish.Idx].Title)
			m.Todo.Delete(m.Todo.Finish.Idx+1, m.todoView)
			m.Todo.Update()
		}
	}
	return m, nil
}

func ctrlt(m *Model) (tea.Model, tea.Cmd) {
	if m.todoActive && m.todoView != 0 {
		if m.todoView == 1 {
			m.Todo.AddTodo(m.Todo.Progress.Title[m.Todo.Progress.Idx].Title, m.Todo.Progress.Desc[m.Todo.Progress.Idx].Title)
			m.Todo.Delete(m.Todo.Progress.Idx+1, m.todoView)
			m.Todo.Update()
		} else if m.todoView == 2 {
			m.Todo.AddTodo(m.Todo.Finish.Title[m.Todo.Finish.Idx].Title, m.Todo.Finish.Desc[m.Todo.Finish.Idx].Title)
			m.Todo.Delete(m.Todo.Finish.Idx+1, m.todoView)
			m.Todo.Update()
		}
	}
	return m, nil
}

func ctrlf(m *Model) (tea.Model, tea.Cmd) {
	if m.todoActive && m.todoView != 2 {
		if m.todoView == 0 {
			m.Todo.AddFinish(m.Todo.Todo.Title[m.Todo.Todo.Idx].Title, m.Todo.Todo.Desc[m.Todo.Todo.Idx].Title)
			m.Todo.Delete(m.Todo.Todo.Idx+1, m.todoView)
			m.Todo.Update()
		} else if m.todoView == 1 {
			m.Todo.AddFinish(m.Todo.Progress.Title[m.Todo.Progress.Idx].Title, m.Todo.Progress.Desc[m.Todo.Progress.Idx].Title)
			m.Todo.Delete(m.Todo.Progress.Idx+1, m.todoView)
			m.Todo.Update()
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
	} else if m.todoActive && m.DeletePopup {
		m.DeletePopup = false
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
	if m.DeletePopup && !m.todoActive {
		m.projectList.Delete(m.projectList.index)
		m.DeletePopup = false
		return m, nil
	} else if m.DeletePopup && m.todoActive {
		if m.todoView == 0 {
			m.DeletePopup = false
			m.Todo.Delete(m.Todo.Todo.Idx+1, m.todoView)
		} else if m.todoView == 1 {
			m.DeletePopup = false
			m.Todo.Delete(m.Todo.Progress.Idx+1, m.todoView)
		} else if m.todoView == 2 {
			m.DeletePopup = false
			m.Todo.Delete(m.Todo.Finish.Idx+1, m.todoView)
		}
		m.Todo.Update()
		return m, nil
	}
	if m.projectActive {
		m.todoActive = true
		m.projectActive = false
		m.Todo.GetTodo(m.projectList.list[m.projectList.index])
		return m, nil
	}
	if m.PopTodo.textareaActive && m.PopTodo.inputActive && m.PopTodo.confirm == 0 {
		m.PopTodo.inputActive = false
		m.PopTodo.textActive = true
		m.PopTodo.input.TextStyle = TextStyleInput
		m.PopTodo.input.SetCursor(0)
		m.PopTodo.input.CursorStyle = TextStyleInput
		m.PopTodo.textarea.Cursor.Style = cursorStyle
		m.PopTodo.textarea.Focus()
		m.PopTodo.textarea.Cursor.Blur()
		m.PopTodo.textarea.Placeholder = " "
		textarea.Blink()
		m.PopTodo.confirm = 1
	} else if m.PopTodo.textareaActive && m.PopTodo.textActive && m.PopTodo.confirm == 1 {
		tmp := m.PopTodo.textarea.Value()
		m.PopTodo.textarea.SetValue(tmp + "\n")
	} else if m.PopTodo.textareaActive && m.PopTodo.confirm == 2 {
		m.PopTodo.inputActive = false
		m.PopTodo.textareaActive = false
		if m.todoView == 0 && m.PopTodo.input.Value() != "" {
			m.Todo.AddTodo(m.PopTodo.inputmsg, m.PopTodo.textarea.Value())
		} else if m.todoView == 1 && m.PopTodo.input.Value() != "" {
			m.Todo.AddProgress(m.PopTodo.inputmsg, m.PopTodo.textarea.Value())
		} else if m.todoView == 2 && m.PopTodo.input.Value() != "" {
			m.Todo.AddFinish(m.PopTodo.inputmsg, m.PopTodo.textarea.Value())
		}
		m.Todo.Update()
	} else if m.PopTodo.textareaActive && m.PopTodo.confirm == 3 {
		m.PopTodo.inputActive = false
		m.PopTodo.textareaActive = false
	}
	return m, nil
}

func down(m *Model) (tea.Model, tea.Cmd) {
	if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
		if m.projectList.index >= m.projectList.size-1 {
			m.projectList.index = 0
		} else {
			m.projectList.index++
		}
	} else if m.todoActive && m.todoView == 0 && m.Todo.Todo.Len > 1 {
		if m.Todo.Todo.Idx == m.Todo.Todo.Len-1 {
			m.Todo.Todo.Idx = 0
		} else {
			m.Todo.Todo.Idx++
		}
	} else if m.todoActive && m.todoView == 1 && m.Todo.Progress.Len > 1 {
		if m.Todo.Progress.Idx == m.Todo.Progress.Len-1 {
			m.Todo.Progress.Idx = 0
		} else {
			m.Todo.Progress.Idx++
		}
	} else if m.todoActive && m.todoView == 2 && m.Todo.Finish.Len > 1 {
		if m.Todo.Finish.Idx == m.Todo.Finish.Len-1 {
			m.Todo.Finish.Idx = 0
		} else {
			m.Todo.Finish.Idx++
		}
	}
	return m, nil
}

func up(m *Model) (tea.Model, tea.Cmd) {
	if m.projectActive && !m.projectAdd && !m.projectRename && !m.typing {
		if m.projectList.index == 0 {
			m.projectList.index = m.projectList.size - 1
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
	if !m.projectAdd && m.projectActive {
		m.addbuffer.Blink()
		m.projectAdd = true
	} else if !m.PopTodo.textareaActive && m.todoActive {
		m.PopTodo.input.Reset()
		m.PopTodo.textareaActive = true
		m.PopTodo.inputActive = true
		m.PopTodo.textActive = false
		m.PopTodo.textarea.Reset()
		m.PopTodo.input.Focus()
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
	} else if m.todoActive {
		m.DeletePopup = true
	}
	m.Todo.Update()
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

func ActiveInput(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.PopTodo.input, cmd = m.PopTodo.input.Update(msg)
	if m.PopTodo.input.Value() != "" {
		m.PopTodo.inputmsg = m.PopTodo.input.Value()
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
