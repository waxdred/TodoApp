package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func leninfo(t []info) int {
	return len(t)
}

func (t *Todo) Addvalue(title, desc string, pos int) {
	if pos == 0 {
	} else if pos == 1 {
	} else if pos == 2 {
	}
}

func (t *Todo) AddTodo(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Todo.Title) + 1,
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Todo.Title) + 1,
	}
	dates := info{
		Title: date,
		Idx:   len(t.Todo.Title) + 1,
	}
	(*t).Todo.Title = append((*t).Todo.Title, titles)
	(*t).Todo.Desc = append((*t).Todo.Desc, descs)
	(*t).Todo.Date = append((*t).Todo.Date, dates)
	if (*t).Todo.Len == -1 {
		(*t).Todo.Len = 1
	} else {
		(*t).Todo.Len++
	}
}

func (t *Todo) AddProgress(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Progress.Title) + 1,
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Progress.Title) + 1,
	}
	dates := info{
		Title: date,
		Idx:   len(t.Progress.Title) + 1,
	}

	(*t).Progress.Title = append((*t).Progress.Title, titles)
	(*t).Progress.Desc = append((*t).Progress.Desc, descs)
	(*t).Progress.Date = append((*t).Todo.Date, dates)
	if (*t).Progress.Len == -1 {
		(*t).Progress.Len = 1
	} else {
		(*t).Progress.Len++
	}
}

func (t *Todo) AddFinish(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Finish.Title) + 1,
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Finish.Title) + 1,
	}
	dates := info{
		Title: date,
		Idx:   len(t.Finish.Title) + 1,
	}
	(*t).Finish.Title = append((*t).Finish.Title, titles)
	(*t).Finish.Desc = append((*t).Finish.Desc, descs)
	(*t).Finish.Date = append((*t).Todo.Date, dates)
	if (*t).Finish.Len == -1 {
		(*t).Finish.Len = 1
	} else {
		(*t).Finish.Len++
	}
}

// TODO need for on the Delete Method need work on because is segfautl
func (t *Todo) Delete(idx, pos int) {
	var todo Todo
	if pos == 0 {
		for i := 0; i < len((*t).Todo.Title); i++ {
			if (*t).Todo.Title[i].Idx != idx {
				todo.AddTodo((*t).Todo.Title[i].Title, (*t).Todo.Desc[i].Title)
			}
		}
		(*t).Todo = todo.Todo
	} else if pos == 1 {
		for i := 0; i < len((*t).Progress.Title); i++ {
			if (*t).Progress.Title[i].Idx != idx {
				todo.AddProgress((*t).Progress.Title[i].Title, (*t).Progress.Desc[i].Title)
			}
		}
		(*t).Progress = todo.Progress
	} else if pos == 2 {
		for i := 0; i < len((*t).Finish.Title); i++ {
			if (*t).Finish.Title[i].Idx != idx {
				todo.AddFinish((*t).Finish.Title[i].Title, (*t).Finish.Desc[i].Title)
			}
		}
		(*t).Finish = todo.Finish
	}
}

func (t *Todo) AddProject(name string) {
	(*t).Project = name
}

func (t *Todo) Update() {
	path := fmt.Sprint(Path, t.Project, ".json")
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(path, file, 0644)
}

func (t *Todo) GetTodo(name string) {
	path := fmt.Sprint(Path, name, ".json")
	_, err := os.Stat(path)
	if err == nil {
		body, errBody := ioutil.ReadFile(path)
		if errBody != nil {
			fmt.Println(errBody)
		}
		jsonErr := json.Unmarshal(body, &t)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	}
}
