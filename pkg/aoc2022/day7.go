package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
)

type Day7 struct {
}

func (d Day7) processInput(input []string) []string {
	return nil
}

func (d Day7) Run() {
	fileName := "input_2022/day7.txt"
	input, err := util.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// inp := d.processInput(input)
	fmt.Println(d.part1(input))
	fmt.Println(d.part2(input))

}

func (d Day7) part1(s string) int {
	return 0
}

func (d Day7) part2(s string) int {
	return 0
}
