package day

import (
	"fmt"
)

type Day3 struct {
}

type sack struct {
	Compartment1 map[rune]int
	Compartment2 map[rune]int
}

func (s *sack) share() rune {
	for k, _ := range s.Compartment1 {
		if _, ok := s.Compartment2[k]; ok {
			return k
		}
	}
	return '@'
}

func priority(r rune) int {
	i := int(r)
	if i <= 90 {
		return i - 38
	}
	return i - 96
}

func toSet(s string) map[rune]int {
	set := make(map[rune]int)
	for _, c := range s {
		if val, ok := set[c]; ok {
			set[c] = val + 1
		}
		set[c] = 1
	}
	return set
}

func (d Day3) processInput(input []string) []*sack {
	sacks := make([]*sack, len(input))
	for i, s := range input {
		l := len(s)
		sacks[i] = &sack{
			Compartment1: toSet(s[0:(l / 2)]),
			Compartment2: toSet(s[(l / 2):]),
		}
	}
	return sacks
}

func (d Day3) Run() {
	var fileName = "input/day3.txt"
	input, err := ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.part1(inp))
	fmt.Println(d.part2(inp))
}

func (d Day3) part1(stacks []*sack) int {
	p := 0
	for _, s := range stacks {
		r := s.share()
		p += priority(r)
	}
	return p
}

func (d Day3) part2(stacks []*sack) int {
	return 0
}
