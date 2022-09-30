package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func ProjectActive(m *Model, row, header, ret string) string {
	var k int
	gap := tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)+(width))))
	header = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	pagn := m.projectList.SizePag(SizeList(row))
	var col string
	if m.typing {
		col = lipgloss.JoinVertical(lipgloss.Top, row, m.search.View())
	} else if m.projectAdd {
		col = lipgloss.JoinVertical(lipgloss.Top, row, m.addbuffer.View())
	} else if m.projectRename {
		rename := fmt.Sprint("Rename: ", m.projectList.list[m.projectList.index])
		m.renamebuffer.Placeholder = rename
		col = lipgloss.JoinVertical(lipgloss.Top, row, m.renamebuffer.View())
	} else {
		col = lipgloss.JoinVertical(lipgloss.Top, row, "\n")
	}
	if m.projectList.index >= SizeList(row) {
		k = SizeList(row)
	} else {
		k = 0
	}
	for i := 0; i < SizeList(row); i++ {
		var matched bool
		if len(m.searchValue) < 2 {
			matched = true
		} else {
			matched, _ = regexp.MatchString(m.searchValue, m.projectList.list[i])
		}
		if lipgloss.Height(col) < height-6 && k < m.projectList.size {
			if k == m.projectList.index && matched {
				col = lipgloss.JoinVertical(lipgloss.Top, col, Selectlist.Render(" > "+m.projectList.list[k]))
			} else if k < m.projectList.size && matched {
				col = lipgloss.JoinVertical(lipgloss.Top, col, NoSelectlist.Render("   "+m.projectList.list[k]))
			}
			k++
		}
	}
	for lipgloss.Height(col) < height-6 {
		col = lipgloss.JoinVertical(lipgloss.Top, col, "\n\n")
	}
	col = ActivatePagn(pagn, SizeList(row), m.projectList.index, col)
	ret = fmt.Sprintf("%s", col)
	return ret
}

func ViewTabTodo(list, progress, finish []string, idx, pos int) string {
	var ret string
	var row string
	for i := 0; i < len(list); i++ {
		if i == 0 && i != idx {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectTop.Render(taskNoSelect.Render("  "+list[i])),
				tabTodoSelectTop.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectTop.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else if i == 0 && i == idx {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectTop.Render(taskNoSelect.Render("> "+list[i])),
				tabTodoSelectTop.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectTop.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else if i == len(list)-1 && i != idx && lipgloss.Height(ret) >= height/4 {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectBottom.Render(taskNoSelect.Render("  "+list[i])),
				tabTodoSelectBottom.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectBottom.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else if i == len(list)-1 && i == idx && lipgloss.Height(ret) >= height/4 {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectBottom.Render(taskNoSelect.Render("> "+list[i])),
				tabTodoSelectBottom.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectBottom.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else if i == idx && lipgloss.Height(ret) >= height/4 {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectMiddle.Render(taskNoSelect.Render("> "+list[i])),
				tabTodoSelectMiddle.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectMiddle.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Top,
				tabTodoSelectMiddle.Render(taskNoSelect.Render("  "+list[i])),
				tabTodoSelectMiddle.Render(taskNoSelect.Render("  "+progress[i])),
				tabTodoSelectMiddle.Render(taskNoSelect.Render("  "+finish[i])))
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		}
	}
	for lipgloss.Height(ret) < height/2 {
		tab := tabTodoSelectMiddle.Render("\n\n")
		row = lipgloss.JoinHorizontal(lipgloss.Top, tab, tab, tab)
		ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
	}
	tab := tabTodoSelectBottom.Render("\n\n")
	row = lipgloss.JoinHorizontal(lipgloss.Top, tab, tab, tab)
	ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
	return ret
}

func ViewDesc(todo interface{}) string {
	todolist := todo.(Todolist)
	desc := lipgloss.JoinVertical(lipgloss.Top, DescriptionSelectTop.Render(DescTiltleStyle.Render(todolist.title)))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescTiltleStyle.Render("Description:")))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescStyle.Render(todolist.desc)))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescTiltleStyle.Render("Creation date:")))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescStyle.Render(todolist.date)))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectBottom.Render(" "))
	return desc
}

func todoView(m *Model, ret string) string {
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabStyle.Render("Project"),
		activeTabStyle.Render(m.spin.View()+"Todo"),
	)
	gap := tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)+(width))))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	// TODO add fonction for chose style
	Task := tabTodoStyleActive.Render("Task")
	Progress := tabTodoStyle.Render("Progress")
	Finish := tabTodoStyle.Render("Finish")
	gap = lipgloss.JoinHorizontal(lipgloss.Top, Task, Progress, Finish)
	row = lipgloss.JoinVertical(lipgloss.Top, row, gap)
	// place Task
	list := []string{
		"test",
		"test0",
		"test1",
	}
	desc := Todolist{
		title: "inception",
		desc:  "Need work on the program add \nnew feat",
		date:  "vendredi 3 novembre",
	}
	task := ViewTabTodo(list, list, list, 0, 0)

	row = lipgloss.JoinVertical(lipgloss.Top, row, task, ViewDesc(desc))
	ret = fmt.Sprintf("%s", row)
	return ret
}
