package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type project struct {
	list      []string
	paginator []string
	index     int
	size      int
}

var user = os.Getenv("USER")

var Path = fmt.Sprint("/Users/", user, "/.config/todo/")

func (m Model) GetSelector(title info) bool {
	if m.todoView == 0 && title.Idx == m.Todo.Todo.Idx+1 {
		return true
	} else if m.todoView == 1 && title.Idx == m.Todo.Progress.Idx+1 {
		return true
	} else if m.todoView == 2 && title.Idx == m.Todo.Finish.Idx+1 {
		return true
	}
	return false
}

func (m Model) GetTilteDesc() string {
	if m.todoView == 0 {
		if m.Todo.Todo.Idx > m.Todo.Todo.Len || m.Todo.Todo.Idx < 0 {
			return ""
		}
		return m.Todo.Todo.Title[m.Todo.Todo.Idx].Title
	} else if m.todoView == 1 {
		if m.Todo.Progress.Idx > m.Todo.Progress.Len || m.Todo.Progress.Idx < 0 {
			return ""
		}
		return m.Todo.Progress.Title[m.Todo.Progress.Idx].Title
	} else if m.todoView == 2 {
		if m.Todo.Finish.Idx > m.Todo.Finish.Len || m.Todo.Finish.Idx < 0 {
			return ""
		}
		return m.Todo.Finish.Title[m.Todo.Finish.Idx].Title
	}
	return ""
}

func (m Model) GetDate() string {
	if m.todoView == 0 {
		if m.Todo.Todo.Idx > m.Todo.Todo.Len || m.Todo.Todo.Idx < 0 {
			return ""
		}
		return m.Todo.Todo.Date[m.Todo.Todo.Idx].Title
	} else if m.todoView == 1 {
		if m.Todo.Progress.Idx > m.Todo.Progress.Len || m.Todo.Progress.Idx < 0 {
			return ""
		}
		return m.Todo.Progress.Date[m.Todo.Progress.Idx].Title
	} else if m.todoView == 2 {
		if m.Todo.Finish.Idx > m.Todo.Finish.Len || m.Todo.Finish.Idx < 0 {
			return ""
		}
		return m.Todo.Finish.Date[m.Todo.Finish.Idx].Title
	}
	return ""
}

func (m Model) GetDesc() string {
	if m.todoView == 0 {
		if m.Todo.Todo.Idx > m.Todo.Todo.Len || m.Todo.Todo.Idx < 1 {
			return ""
		}
		return m.Todo.Todo.Desc[m.Todo.Todo.Idx].Title
	} else if m.todoView == 1 {
		if m.Todo.Progress.Idx > m.Todo.Progress.Len || m.Todo.Progress.Idx < 1 {
			return ""
		}
		return m.Todo.Progress.Desc[m.Todo.Progress.Idx].Title
	} else if m.todoView == 2 {
		if m.Todo.Finish.Idx > m.Todo.Finish.Len || m.Todo.Finish.Idx < 1 {
			return ""
		}
		return m.Todo.Finish.Desc[m.Todo.Finish.Idx].Title
	}
	return ""
}

func (t *project) GetProjectSave() {
	var file string
	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		os.Mkdir(Path, os.ModePerm)
	} else {
		files, errFiles := ioutil.ReadDir(Path)
		if errFiles != nil {
			return
		}
		for _, f := range files {
			if strings.Index(f.Name(), ".json") > -1 {
				file = strings.Replace(f.Name(), ".json", "", 1)
				t.Add(file)
			}
		}
	}
}

func InitTodo(title string) Todo {
	var todo Todo
	todo.Project = title
	todo.Todo.Len = -1
	todo.Progress.Len = -1
	todo.Finish.Len = -1
	return todo
}

func (t *project) Add(title string) bool {
	var todo Todo
	title = strings.Replace(title, " ", "_", -1)
	path := fmt.Sprint(Path, title, ".json")
	todo.Project = title
	todo.Todo.Len = -1
	todo.Progress.Len = -1
	todo.Finish.Len = -1
	if !t.CmpTittle(title) {
		return false
	}
	if title != "" {
		(*t).list = append((*t).list, title)
		(*t).size++
		(*t).index = 0
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			file, _ := json.MarshalIndent(&todo, "", " ")
			ioutil.WriteFile(path, file, 0644)
		}
	}
	return true
}

func (t project) CmpTittle(title string) bool {
	for _, lst := range t.list {
		if lst == title {
			return false
		}
	}
	return true
}

func InitProjectList() *project {
	t := new(project)

	(*t).index = 0
	(*t).size = 0
	t.GetProjectSave()
	return t
}

func SizeList(t string) int {
	sizeT := lipgloss.Height(t)
	ret := height - sizeT - 6
	return ret / 3
}

func (t project) SizePag(size int) int {
	val := float64(float64(t.size) / float64(size))
	valInt := int(t.size / size)
	if val > float64(valInt) {
		return valInt + 1
	}
	return valInt
}

func (t *project) Rename(title string, idx int) {
	if (*t).list[idx] != "" {
		title = strings.Replace(title, " ", "_", -1)
		old := fmt.Sprint(Path, (*t).list[idx], ".json")
		replace := fmt.Sprint(Path, title, ".json")
		os.Rename(old, replace)
		(*t).list[idx] = title
	}
}

func (t *project) Delete(idx int) {
	var tmp []string
	path := fmt.Sprint(Path, (*t).list[idx], ".json")
	i := 0
	if idx <= (*t).size {
		for _, lst := range (*t).list {
			if i != idx {
				tmp = append(tmp, lst)
			}
			i++
		}
		(*t).list = tmp
	}
	(*t).size--
	if (*t).index > (*t).size {
		(*t).index = (*t).size
	}
	os.Remove(path)
}
