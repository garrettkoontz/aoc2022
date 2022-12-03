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

func (s *sack) combined() map[rune]int {
	newMap := make(map[rune]int)
	for k, v := range s.Compartment1 {
		newMap[k] = v
	}
	for k, v2 := range s.Compartment2 {
		v, ok := s.Compartment2[k]
		if ok {
			newMap[k] = v + v2
			continue
		}
		newMap[k] = v2
	}
	return newMap
}

func share(i1 map[rune]int, i2 map[rune]int, i3 map[rune]int) rune {
	share1 := make(map[rune]bool)
	for k, _ := range i1 {
		if _, ok := i2[k]; ok {
			share1[k] = true
		}
	}
	for k, _ := range share1 {
		if _, ok := i3[k]; ok {
			return k
		}
	}
	return '@'
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
	p := 0
	for i := 0; i < len(stacks); i += 3 {
		c := share(stacks[i].combined(), stacks[i+1].combined(), stacks[i+2].combined())
		p += priority(c)
	}
	return p
}
