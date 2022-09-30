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

type Todolist struct {
	Title string
	Desc  string
	Date  string
}

type Todo struct {
	Project string `json:"project"`
	Todo    struct {
		Title []string `json:"title"`
		Desc  []string `json:"desc"`
		Date  []string `json:"date"`
		Idx   int      `json:"idx"`
		Len   int      `json:"len"`
	} `json:"todo"`
	Progress struct {
		Title []string `json:"title"`
		Desc  []string `json:"desc"`
		Date  []string `json:"date"`
		Idx   int      `json:"idx"`
		Len   int      `json:"len"`
	} `json:"progress"`
	Finish struct {
		Title []string `json:"title"`
		Desc  []string `json:"desc"`
		Date  []string `json:"date"`
		Idx   int      `json:"idx"`
		Len   int      `json:"len"`
	} `json:"finish"`
}
