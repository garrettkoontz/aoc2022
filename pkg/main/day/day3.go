package day

import (
	"fmt"
)

type Day3 struct {
}

func (d Day3) processInput(input []string) int {
	return 0
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

func (d Day3) part1(i int) int {
	return 0
}

func (d Day3) part2(i int) int {
	return 0
}
