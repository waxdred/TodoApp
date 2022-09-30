package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func (t *Todo) AddTodo(title, desc string) {
	ts := time.Now()
	date := fmt.Sprint(ts.Year, " ", ts.Month, " ", ts.Day)
	(*t).Todo.Title = append((*t).Todo.Title, title)
	(*t).Todo.Desc = append((*t).Todo.Desc, title)
	(*t).Todo.Date = append((*t).Todo.Date, date)
	(*t).Todo.Len++
}

func (t *Todo) AddProgress(title, desc string) {
	ts := time.Now()
	date := fmt.Sprint(ts.Year, " ", ts.Month, " ", ts.Day)
	(*t).Progress.Title = append((*t).Progress.Title, title)
	(*t).Progress.Desc = append((*t).Progress.Desc, title)
	(*t).Progress.Date = append((*t).Todo.Date, date)
	(*t).Progress.Len++
}

func (t *Todo) AddFinish(title, desc string) {
	ts := time.Now()
	date := fmt.Sprint(ts.Year, " ", ts.Month, " ", ts.Day)
	(*t).Finish.Title = append((*t).Finish.Title, title)
	(*t).Finish.Desc = append((*t).Finish.Desc, title)
	(*t).Finish.Date = append((*t).Todo.Date, date)
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
		jsonErr := json.Unmarshal(body, t)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	}
}
