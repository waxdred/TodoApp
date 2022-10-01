package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Todo           Todo
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
	AddPopup       bool
	todoView       int
	err            error
}

type info struct {
	Title string `json:"title"`
	Idx   int    `json:"idx"`
}

type Todolist struct {
	Title []info `json:"title"`
	Desc  []info `json:"desc"`
	Date  []info `json:"date"`
	Idx   int    `json:"idx"`
	Len   int    `json:"len"`
}

type Todo struct {
	Project  string   `json:"project"`
	Todo     Todolist `json:"todo"`
	Progress Todolist `json:"progress"`
	Finish   Todolist `json:"finish"`
}
