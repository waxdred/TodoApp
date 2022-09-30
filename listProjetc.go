package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type project struct {
	list      []string
	paginator []string
	index     int
	size      int
}

var user = os.Getenv("USER")

var Path = fmt.Sprint("/Users/", user, "/.config/todo/")

func (t *project) GetProjectSave() {
	var file string
	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		os.Mkdir(Path, os.ModePerm)
	} else {
		files, errFiles := ioutil.ReadDir(Path)
		if errFiles != nil {
			return
		}
		for _, f := range files {
			if strings.Index(f.Name(), ".json") > -1 {
				file = strings.Replace(f.Name(), ".json", "", 1)
				t.Add(file)
			}
		}
	}
}

func (t *project) Add(title string) bool {
	path := fmt.Sprint(Path, title, ".json")
	if !t.CmpTittle(title) {
		return false
	}
	if title != "" {
		title = strings.Replace(title, " ", "_", -1)
		(*t).list = append((*t).list, title)
		(*t).size++
		(*t).index = 0
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			os.Create(path)
		}
	}
	return true
}

func (t project) CmpTittle(title string) bool {
	for _, lst := range t.list {
		if lst == title {
			return false
		}
	}
	return true
}

func InitProjectList() *project {
	t := new(project)

	(*t).index = 0
	(*t).size = 0
	t.GetProjectSave()
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

func (t *project) Rename(title string, idx int) {
	if (*t).list[idx] != "" {
		title = strings.Replace(title, " ", "_", -1)
		old := fmt.Sprint(Path, (*t).list[idx], ".json")
		replace := fmt.Sprint(Path, title, ".json")
		os.Rename(old, replace)
		(*t).list[idx] = title
	}
}

func (t *project) Delete(idx int) {
	var tmp []string
	path := fmt.Sprint(Path, (*t).list[idx], ".json")
	i := 0
	if idx <= (*t).size {
		for _, lst := range (*t).list {
			if i != idx {
				tmp = append(tmp, lst)
			}
			i++
		}
		(*t).list = tmp
	}
	(*t).size--
	if (*t).index > (*t).size {
		(*t).index = (*t).size
	}
	os.Remove(path)
}
