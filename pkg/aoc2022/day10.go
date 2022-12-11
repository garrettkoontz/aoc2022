package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"regexp"
	"strconv"
)

type Day10 struct {
}

type CPU struct {
	Count    int
	Register int
}

func (c *CPU) run(ins []*int) []int {
	strength := []int{c.Register}
	for _, addx := range ins {
		if addx == nil {
			c.Count += 1
			strength = append(strength, c.Register)
			continue
		}
		strength = append(strength, c.Register)
		c.Count += 2
		c.Register += *addx
		strength = append(strength, c.Register)
	}
	return strength
}

func draw(vals []int) [][]rune {
	view := make([][]rune, 7)
	for i := range view {
		view[i] = make([]rune, 40)
	}
	for i, v := range vals {
		row := i / 40
		col := i % 40
		if v == col || v-1 == col || v+1 == col {
			view[row][col] = '#'
			continue
		}
		view[row][col] = ' '
	}
	return view
}

func (d Day10) processInput(input []string) []*int {
	result := make([]*int, len(input))
	for i, line := range input {
		r := regexp.MustCompile("addx ([-]?[0-9]+)")
		s := r.FindStringSubmatch(line)
		if s == nil {
			result[i] = nil
			continue
		}
		j, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		result[i] = &j
	}
	return result
}

func (d Day10) Run() {
	fileName := "input_2022/day10.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	d.part2(dir)
}

func (d Day10) part1(instructions []*int) int {
	cpu := CPU{Count: 1, Register: 1}
	strength := cpu.run(instructions)
	return 20*strength[19] +
		60*strength[59] +
		100*strength[99] +
		140*strength[139] +
		180*strength[179] +
		220*strength[219]
}

func (d Day10) part2(instructions []*int) {
	cpu := CPU{Count: 1, Register: 1}
	vals := cpu.run(instructions)
	view := draw(vals)
	print(view)
}

func print(view [][]rune) {
	for _, v := range view {
		for _, c := range v {
			fmt.Printf("%s", string(c))
		}
		fmt.Printf("\n")
	}
}

// too low 5463
