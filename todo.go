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

func (t *Todo) AddTodo(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Todo.Title),
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Todo.Title),
	}
	dates := info{
		Title: date,
		Idx:   len(t.Todo.Title),
	}
	(*t).Todo.Title = append((*t).Todo.Title, titles)
	(*t).Todo.Desc = append((*t).Todo.Desc, descs)
	(*t).Todo.Date = append((*t).Todo.Date, dates)
	(*t).Todo.Len++
}

func (t *Todo) AddProgress(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Progress.Title),
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Progress.Title),
	}
	dates := info{
		Title: date,
		Idx:   len(t.Progress.Title),
	}

	(*t).Progress.Title = append((*t).Progress.Title, titles)
	(*t).Progress.Desc = append((*t).Progress.Desc, descs)
	(*t).Progress.Date = append((*t).Todo.Date, dates)
	(*t).Progress.Len++
}

func (t *Todo) AddFinish(title, desc string) {
	ts := time.Now()
	date := ts.String()
	titles := info{
		Title: title,
		Idx:   len(t.Finish.Title),
	}
	descs := info{
		Title: desc,
		Idx:   len(t.Finish.Title),
	}
	dates := info{
		Title: date,
		Idx:   len(t.Finish.Title),
	}
	(*t).Finish.Title = append((*t).Finish.Title, titles)
	(*t).Finish.Desc = append((*t).Finish.Desc, descs)
	(*t).Finish.Date = append((*t).Todo.Date, dates)
	(*t).Finish.Len++
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
	os.Stdout.Write(file)
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
