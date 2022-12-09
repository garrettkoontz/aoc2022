package aoc2016

import (
	"aoc2022/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct {
}

type direction rune

const (
	l direction = 'L'
	r direction = 'R'
)

type Move struct {
	Direction rune
	Amount    int
}

var cardinal = []util.Position2D{
	{X: 0, Y: 1},  // N
	{X: 1, Y: 0},  // E
	{X: 0, Y: -1}, // S
	{X: -1, Y: 0}, // W
}

func (d *Day1) processInput(input string) []*Move {
	s := strings.Split(input, ", ")
	result := make([]*Move, len(s))
	for i, v := range s {
		m := rune(v[0])
		a, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}
		move := Move{
			Direction: m,
			Amount:    a,
		}
		result[i] = &move
	}
	return result
}

func (d *Day1) Run() {
	fileName := "input_2016/day1.txt"
	input, err := util.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.Part1(inp))
	fmt.Println(d.Part2(inp))
}

func (d *Day1) Part1(moves []*Move) int {
	p := util.Position2D{X: 0, Y: 0}
	dir := 0
	for _, m := range moves {
		d := (dir + 5) % 4
		if m.Direction == 'L' {
			d = (dir + 3) % 4
		}
		turn := &cardinal[d]
		dir = d
		turn.Times(m.Amount)
		p.Plus(turn)
	}
	return p.Distance()
}

func (d *Day1) Part2(moves []*Move) int {
	visits := make(map[util.Position2D]bool)
	p := util.Position2D{X: 0, Y: 0}
	visits[p] = true
	dir := 0
	for _, m := range moves {
		d := (dir + 5) % 4
		if m.Direction == 'L' {
			d = (dir + 3) % 4
		}
		turn := &cardinal[d]
		dir = d
		for i := 0; i < m.Amount; i++ {
			p.Plus(turn)
			if _, ok := visits[p]; ok {
				return p.Distance()
			}
			visits[p] = true
		}
	}
	return -1
}
