package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
}

type Dir struct {
	Name   string
	Parent *Dir
	Files  []*file
	Dirs   []*Dir
	Size   *int
}

func (d *Dir) cd(name string) *Dir {
	if d == nil {
		return &Dir{Name: name}
	}
	if name == ".." {
		return d.Parent
	}
	for _, dr := range d.Dirs {
		if dr.Name == name {
			return dr
		}
	}
	return nil
}

func (d *Dir) AddSubDir(subD *Dir) {
	d.Dirs = append(d.Dirs, subD)
	subD.Parent = d
}

func (d *Dir) GetSize() int {
	if d.Size == nil {
		s := 0
		for _, f := range d.Files {
			s += f.Size
		}
		for _, subD := range d.Dirs {
			s += subD.GetSize()
		}
		d.Size = &s
		return s
	}
	return *d.Size
}

type file struct {
	Size int
	Name string
}

func (d Day7) processInput(input []string) *Dir {
	dirName := ""
	var currDir *Dir
	top := currDir
	for i, line := range input {
		// fmt.Printf("i: %d, line: %q, dir: %+v, dirp:%p \n", i, line, currDir, currDir)
		// command
		if line[0] == '$' {
			if line == "$ ls" {
				continue
			}
			dirName = line[5:]

			currDir = currDir.cd(dirName)

			if i == 0 {
				top = currDir
			}
			continue
		}
		// listing dir
		if line[:3] == "dir" {
			name := line[4:]
			dn := Dir{Name: name, Parent: currDir}
			currDir.AddSubDir(&dn)
			continue
		}
		// file
		s := strings.Split(line, " ")
		size, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		fn := s[1]
		f := file{size, fn}
		currDir.Files = append(currDir.Files, &f)
	}
	return top
}

func (d Day7) Run() {
	fileName := "input_2022/day7.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	fmt.Println(d.part2(dir))
}

func (d Day7) part1(dir *Dir) int {
	dir.GetSize()
	s := 0
	queue := []*Dir{dir}
	for i := 0; i < len(queue); i++ {
		d := queue[i]
		queue = append(queue, d.Dirs...)
		sz := *d.Size
		if sz <= 100000 {
			s += sz
		}
	}
	return s
}

func (d Day7) part2(dir *Dir) int {
	totalAvaiable := 70000000
	neededSpace := 30000000
	currentUnusedSize := totalAvaiable - *dir.Size
	deleteMinSpace := neededSpace - currentUnusedSize
	sizeOfDeleted := *dir.Size
	queue := []*Dir{dir}
	for i := 0; i < len(queue); i++ {
		d := queue[i]
		queue = append(queue, d.Dirs...)
		sz := *d.Size
		if sz > deleteMinSpace && sz < sizeOfDeleted {
			sizeOfDeleted = sz
		}
	}
	return sizeOfDeleted
}
