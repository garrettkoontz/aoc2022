package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day4 struct {
}

type assignmentPair struct {
	A1 map[int]bool
	A2 map[int]bool
}

func (a *assignmentPair) fullyContains() bool {
	oneContainsTwo := true
	for k := range a.A1 {
		if _, ok := a.A2[k]; !ok {
			oneContainsTwo = false
			break
		}
	}

	twoContainsOne := true
	for k := range a.A2 {
		if _, ok := a.A1[k]; !ok {
			twoContainsOne = false
			break
		}
	}
	return oneContainsTwo || twoContainsOne
}

func (a *assignmentPair) overlap() bool {
	for k := range a.A1 {
		if _, ok := a.A2[k]; ok {
			return true
		}
	}
	return false
}

func (d Day4) Run() {
	fileName := "input/day4.txt"
	input, err := ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.part1(inp))
	fmt.Println(d.part2(inp))
}

func (d Day4) part1(ap []*assignmentPair) int {
	contains := 0
	for _, ap2 := range ap {
		if ap2.fullyContains() {
			contains += 1
		}
	}
	return contains
}

func (d Day4) part2(ap []*assignmentPair) int {
	overlaps := 0
	for _, ap2 := range ap {
		if ap2.overlap() {
			overlaps += 1
		}
	}
	return overlaps
}

func toIntSet(sp string) map[int]bool {
	sp1 := strings.Split(sp, "-")
	start, err := strconv.Atoi(sp1[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(sp1[1])
	if err != nil {
		panic(err)
	}
	return makeIntSet(start, end)
}
func makeIntSet(s, e int) map[int]bool {
	mp := make(map[int]bool, e-s)
	for i := s; i <= e; i++ {
		mp[i] = true
	}
	return mp
}

func (d Day4) processInput(input []string) []*assignmentPair {
	result := make([]*assignmentPair, len(input))
	for i, s := range input {
		sp := strings.Split(s, ",")
		a1 := toIntSet(sp[0])
		a2 := toIntSet(sp[1])
		result[i] = &assignmentPair{
			A1: a1,
			A2: a2,
		}
	}
	return result
}

// 291 too low
