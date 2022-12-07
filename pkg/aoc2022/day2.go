package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"strings"
)

type Day2 struct {
}

type RockPaperScissor int

const (
	Rock    RockPaperScissor = 1
	Paper   RockPaperScissor = 2
	Scissor RockPaperScissor = 3
)

func game(me RockPaperScissor, oppo RockPaperScissor) int {
	score := int(me)
	if me == oppo {
		return score + 3
	}
	if int(me)-1 == int(oppo) || (me == Rock && oppo == Scissor) {
		return score + 6
	}
	return score
}

func lose(i int) int {
	if i == 1 {
		return 3
	}
	return i - 1
}

func win(i int) int {
	if i == 3 {
		return 1
	}
	return i + 1
}

func cheat(me RockPaperScissor, oppo RockPaperScissor) int {
	if int(me) == 2 {
		// draw
		return int(oppo) + 3
	}
	if int(me) == 1 {
		// lost
		return lose(int(oppo))
	}
	if int(me) == 3 {
		// win
		return win(int(oppo)) + 6
	}
	return 0
}

func parseRPS(s string) RPS {
	split := strings.Split(s, " ")
	var o RockPaperScissor
	switch split[0] {
	case "A":
		o = Rock
		break
	case "B":
		o = Paper
		break
	case "C":
		o = Scissor
		break
	default:
		panic(s[1])
	}
	var m RockPaperScissor
	switch split[1] {
	case "X":
		m = Rock
		break
	case "Y":
		m = Paper
		break
	case "Z":
		m = Scissor
		break
	default:
		panic(s[1])
	}
	return RPS{
		Opponent: o,
		Me:       m,
	}
}

type RPS struct {
	Opponent RockPaperScissor
	Me       RockPaperScissor
}

func (r *RPS) score() int {
	return game(r.Me, r.Opponent)
}

func (r *RPS) cheat() int {
	return cheat(r.Me, r.Opponent)
}

func (d Day2) processInput(input []string) []*RPS {
	result := []*RPS{}
	for _, g := range input {
		rps := parseRPS(g)
		result = append(result, &rps)
	}
	return result
}

func (d Day2) Run() {
	var fileName = "input_2022/day2.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.part1(inp))
	fmt.Println(d.part2(inp))
}

func (d Day2) part1(rps []*RPS) int {
	s := 0
	for _, g := range rps {
		s += g.score()
	}
	return s
}

func (d Day2) part2(rps []*RPS) int {
	s := 0
	for _, g := range rps {
		s += g.cheat()
	}
	return s
}
