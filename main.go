package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initModel := Model{
		projectList: InitProjectList(),
		spin:        SpinInit(spinStyle),
		search:      InputInit(searchStyle),
		// TODO: need change with touch up or down or /
		typing: false,

		projectActive: true,
		projet:        activeTabStyle,
		todoActive:    false,
		todo:          tabStyle,
		exitPopup:     false,
	}
	initModel.projectList.Add("test")
	initModel.projectList.Add("test1")
	initModel.projectList.Add("test2")
	initModel.projectList.Add("test3")
	initModel.projectList.Add("test4")
	initModel.projectList.Add("test5")
	initModel.projectList.Add("test6")
	initModel.projectList.Add("test7")
	initModel.projectList.Add("test8")
	initModel.projectList.Add("test9")
	initModel.projectList.Add("test10")
	initModel.projectList.Add("test11")
	initModel.projectList.Add("test12")
	initModel.projectList.Add("test13")
	initModel.projectList.Add("test14")
	initModel.projectList.Add("test15")
	initModel.projectList.Add("test16")
	initModel.projectList.Add("test17")
	initModel.projectList.Add("test18")
	initModel.projectList.Add("test19")
	errTea := tea.NewProgram(&initModel, tea.WithAltScreen()).Start()
	if errTea != nil {
		fmt.Println("Error: ", errTea)
		os.Exit(1)
	}
}
