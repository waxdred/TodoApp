package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
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
	AddPopup       bool
	err            error
}
