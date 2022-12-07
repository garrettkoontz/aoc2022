package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
)

type Day6 struct {
}

func (d Day6) processInput(input []string) []string {
	return nil
}

func (d Day6) Run() {
	fileName := "input_2022/day6.txt"
	input, err := util.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// inp := d.processInput(input)
	fmt.Println(d.part1(input))
	fmt.Println(d.part2(input))

}

func (d Day6) part1(s string) int {
	return distinctRunIndex(s, 4)
}

func distinctRunIndex(s string, d int) int {
	for i := 0; i < len(s); i++ {
		mp := make(map[byte]bool)
		for j := i; j < i+d; j++ {
			mp[s[j]] = true
		}
		if len(mp) == d {
			return i + d
		}
	}
	return -1
}

func (d Day6) part2(s string) int {
	return distinctRunIndex(s, 14)
}
