package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"sort"
	"strconv"
)

type Day1 struct {
}

func (d Day1) Run() {
	fileName := "input_2022/day1.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.part1(inp))
	fmt.Println(d.part2(inp))
}

func (d Day1) part1(elves []ElfCalories) int {
	max := 0
	for _, e := range elves {
		tc := e.TotalCalories()
		if tc > max {
			max = tc
		}
	}
	return max
}

func (d Day1) part2(elves []ElfCalories) int {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories() > elves[j].TotalCalories()
	})
	return elves[0].TotalCalories() + elves[1].TotalCalories() + elves[2].TotalCalories()
}

func (e ElfCalories) TotalCalories() int {
	i := 0
	for _, c := range e.Calories {
		i += c
	}
	return i
}

type ElfCalories struct {
	Calories []int
}

func (d Day1) processInput(input []string) []ElfCalories {
	fmt.Println(len(input))
	elfCals := []ElfCalories{}
	cals := []int{}
	for _, s := range input {
		if s == "" {
			elfCals = append(elfCals, ElfCalories{
				Calories: cals,
			},
			)
			cals = []int{}
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		cals = append(cals, i)
	}
	return elfCals
}
