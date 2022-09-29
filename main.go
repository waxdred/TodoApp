package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initModel := Model{
		projectAdd:    false,
		projectRename: false,
		projectList:   InitProjectList(),
		spin:          SpinInit(spinStyle),
		search:        InputInit(searchStyle, " Search "),
		addbuffer:     InputInit(searchStyle, " Add "),
		renamebuffer:  InputInit(searchStyle, " Rename "),
		// TODO: need change with touch up or down or /
		typing: false,

		projectActive: true,
		projet:        activeTabStyle,
		todoActive:    false,
		todo:          tabStyle,
		exitPopup:     false,
		DeletePopup:   false,
		AddPopup:      false,
	}
	errTea := tea.NewProgram(&initModel, tea.WithAltScreen()).Start()
	if errTea != nil {
		fmt.Println("Error: ", errTea)
		os.Exit(1)
	}
}
