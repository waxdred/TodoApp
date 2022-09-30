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

// TODO Work here need found better solution
func ViewTodoRow(Select, noSelect, borderSelect, borderNoSelect lipgloss.Style, pos int, title ...string) string {
	var row string

	for i := 0; i < len(title); i++ {
		if pos == 0 && i == 0 {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, borderSelect.Render(Select.Render(title[i])))
		} else if pos == 1 && i == 1 {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, borderSelect.Render(Select.Render(title[i])))
		} else if pos == 2 && i == 2 {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, borderSelect.Render(Select.Render(title[i])))
		} else {
			row = lipgloss.JoinHorizontal(lipgloss.Top, row, borderNoSelect.Render(noSelect.Render(title[i])))
		}
	}
	return row
}

func checkTitleRow(
	Select, noSelect, borderSelect, borderNoSelect lipgloss.Style,
	pos, i int,
	sizes []int,
	title ...[]string,
) string {
	var row string
	if i < sizes[0] && i < sizes[1] && i < sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  "+title[0][i], "  "+title[1][i], "  "+title[2][i])
		// ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
	} else if i >= sizes[0] && i < sizes[1] && i < sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  ", "  "+title[1][i], "  "+title[2][i])
	} else if i >= sizes[0] && i >= sizes[1] && i < sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  ", "  ", "  "+title[2][i])
	} else if i < sizes[0] && i < sizes[1] && i >= sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  "+title[0][i], "  "+title[1][i], "  ")
	} else if i < sizes[0] && i >= sizes[1] && i >= sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  "+title[0][i], "  ", "  ")
	} else if i >= sizes[0] && i < sizes[1] && i >= sizes[2] {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  ", "  "+title[1][i], "  ")
	} else {
		row = ViewTodoRow(Select,
			noSelect,
			borderSelect,
			borderNoSelect,
			pos, "  ", "  ", "  ")
	}
	return row
}

func CompletedTab(ret string, pos int) string {
	var row1 string
	row := tabTodoNoSelectMiddle.Render(taskNoSelect.Render(" "))
	rowSelect := tabTodoSelectMiddle.Render(taskSelect.Render(" "))
	rowB := tabTodoNoSelectBottom.Render(taskNoSelect.Render(" "))
	rowBSelect := tabTodoSelectBottom.Render(taskSelect.Render(" "))
	for i := 0; lipgloss.Height(ret) <= height/2; i++ {
		if pos == 0 {
			row1 = lipgloss.JoinHorizontal(lipgloss.Top, rowSelect, row, row)
		} else if pos == 1 {
			row1 = lipgloss.JoinHorizontal(lipgloss.Top, row, rowSelect, row)
		} else if pos == 2 {
			row1 = lipgloss.JoinHorizontal(lipgloss.Top, row, row, rowSelect)
		}
		ret = lipgloss.JoinVertical(lipgloss.Top, ret, row1)
	}
	if pos == 0 {
		row1 = lipgloss.JoinHorizontal(lipgloss.Top, rowBSelect, rowB, rowB)
	} else if pos == 1 {
		row1 = lipgloss.JoinHorizontal(lipgloss.Top, rowB, rowBSelect, rowB)
	} else if pos == 2 {
		row1 = lipgloss.JoinHorizontal(lipgloss.Top, rowB, rowB, rowBSelect)
	}
	ret = lipgloss.JoinVertical(lipgloss.Top, ret, row1)
	return ret
}

func ViewTabTodo(list, progress, finish []string, pos int) string {
	var ret string
	var row string
	sizes := []int{
		len(list),
		len(progress),
		len(finish),
	}
	size := max(sizes[0], sizes[1])
	size = max(size, sizes[2])
	if size == 0 {
		row = checkTitleRow(taskSelect,
			taskNoSelect,
			tabTodoSelectTop,
			tabTodoNoSelectTop,
			pos, 0, sizes, list, progress, finish)
		ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
	}
	for i := 0; i < size; i++ {
		if i == 0 {
			row = checkTitleRow(taskSelect,
				taskNoSelect,
				tabTodoSelectTop,
				tabTodoNoSelectTop,
				pos, i, sizes, list, progress, finish)
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		} else if lipgloss.Height(ret) <= height/6 {
			row = checkTitleRow(taskSelect,
				taskNoSelect,
				tabTodoSelectMiddle,
				tabTodoNoSelectMiddle,
				pos, i, sizes, list, progress, finish)
			ret = lipgloss.JoinVertical(lipgloss.Top, ret, row)
		}
	}
	ret = CompletedTab(ret, pos)
	return ret
}

// TODO new work on
func ViewDesc(todo interface{}, idx int) string {
	todolist := todo.(Todolist)
	desc := lipgloss.JoinVertical(lipgloss.Top, DescriptionSelectTop.Render(DescTiltleStyle.Render(todolist.Title[idx])))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescNorStyle.Render("Description:")))
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(DescStyle.Render(todolist.Desc[idx])))
	desc = lipgloss.JoinVertical(
		lipgloss.Top,
		desc,
		DescriptionSelectMiddle.Render(DescNorStyle.Render("Creation date:")+DescDateStyle.Render(todolist.Date[idx])),
	)
	for lipgloss.Height(desc) <= height/4 {
		desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectMiddle.Render(" "))
	}
	desc = lipgloss.JoinVertical(lipgloss.Top, desc, DescriptionSelectBottom.Render(" "))
	return desc
}

func ViewTask(idx int) []string {
	var ret []string
	if idx == 0 {
		ret = append(ret, tabTodoSelect.Render("Task"))
	} else {
		ret = append(ret, tabTodoNoSelect.Render("Task"))
	}
	if idx == 1 {
		ret = append(ret, tabTodoSelect.Render("Progress"))
	} else {
		ret = append(ret, tabTodoNoSelect.Render("Progress"))
	}
	if idx == 2 {
		ret = append(ret, tabTodoSelect.Render("Finish"))
	} else {
		ret = append(ret, tabTodoNoSelect.Render("Finish"))
	}
	return ret
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
	task := ViewTask(m.todoView)
	gap = lipgloss.JoinHorizontal(lipgloss.Top, task[0], task[1], task[2])
	row = lipgloss.JoinVertical(lipgloss.Top, row, gap)
	// place Task
	tasks := ViewTabTodo(m.Todo.Todo.Title, m.Todo.Progress.Title, m.Todo.Finish.Title, m.todoView)

	row = lipgloss.JoinVertical(lipgloss.Top, row, tasks)
	desc := ViewDesc(m.Todo.Todo, 0)
	row = lipgloss.JoinVertical(lipgloss.Top, row, desc)
	helper := fmt.Sprint(
		"\nAdd: <ctrl+a>   Modify: <ctrl+r>   Delete: <ctrl+d>   nav: arrow  Back to Project: <Esc>",
	)
	row = lipgloss.JoinVertical(lipgloss.Top, row, helper)
	ret = fmt.Sprintf("%s", row)
	return ret
}
