package main

import (
	"github.com/charmbracelet/lipgloss"
)

type project struct {
	list      []string
	paginator []string
	index     int
	size      int
}

func (t *project) Add(title string) {
	(*t).list = append((*t).list, title)
	(*t).size++
}

func InitProjectList() *project {
	t := new(project)

	(*t).index = 0
	(*t).size = 0
	return t
}

func SizeList(t string) int {
	sizeT := lipgloss.Height(t)
	ret := height - sizeT - 6
	return ret / 3
}

func (t project) SizePag(size int) int {
	val := float64(float64(t.size) / float64(size))
	valInt := int(t.size / size)
	if val > float64(valInt) {
		return valInt + 1
	}
	return valInt
}
