package aoc2016

import (
	"aoc2022/pkg/util"
	"fmt"
)

type Day2 struct {
}

var (
	up    util.Position2D = util.Position2D{X: 0, Y: -1}
	down  util.Position2D = util.Position2D{X: 0, Y: 1}
	left  util.Position2D = util.Position2D{X: -1, Y: 0}
	right util.Position2D = util.Position2D{X: 1, Y: 0}
)

var (
	one   util.Position2D = util.Position2D{X: -1, Y: -1}
	two   util.Position2D = util.Position2D{X: 0, Y: -1}
	three util.Position2D = util.Position2D{X: 1, Y: -1}
	four  util.Position2D = util.Position2D{X: -1, Y: 0}
	five  util.Position2D = util.Position2D{X: 0, Y: 0}
	six   util.Position2D = util.Position2D{X: 1, Y: 0}
	seven util.Position2D = util.Position2D{X: -1, Y: 1}
	eight util.Position2D = util.Position2D{X: 0, Y: 1}
	nine  util.Position2D = util.Position2D{X: 1, Y: 1}
)

func move1(position *util.Position2D, move *util.Position2D) {
	if util.Abs(position.X+move.X) <= 1 && util.Abs(position.Y+move.Y) <= 1 {
		position.Plus(move)
	}
}

func (d *Day2) processInput(input []string) [][]*util.Position2D {
	result := make([][]*util.Position2D, len(input))
	for i, inp := range input {
		res := make([]*util.Position2D, len(inp))
		for j, r := range inp {
			switch r {
			case 'U':
				res[j] = &up
				break
			case 'D':
				res[j] = &down
				break
			case 'L':
				res[j] = &left
				break
			case 'R':
				res[j] = &right
				break
			default:
				panic(r)
			}
		}
		result[i] = res
	}
	return result
}

func (d *Day2) Run() {
	fileName := "input_2016/day2.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.Part1(inp))
	fmt.Println(d.Part2(inp))
}

func (d *Day2) Part1(moves [][]*util.Position2D) string {
	str := []rune{}
	start := &util.Position2D{X: 0, Y: 0}
	for _, v := range moves {
		for _, q := range v {
			move1(start, q)
		}
		s := *start
		switch s {
		case one:
			str = append(str, '1')
			break
		case two:
			str = append(str, '2')
			break
		case three:
			str = append(str, '3')
			break
		case four:
			str = append(str, '4')
			break
		case five:
			str = append(str, '5')
			break
		case six:
			str = append(str, '6')
			break
		case seven:
			str = append(str, '7')
			break
		case eight:
			str = append(str, '8')
			break
		case nine:
			str = append(str, '9')
			break
		default:
			panic(s)
		}

	}

	return string(str)
}

func move(p *util.Position2D, move *util.Position2D, mp map[util.Position2D]rune) {
	newP := util.Position2D{X: p.X + move.X, Y: p.Y + move.Y}
	if _, ok := mp[newP]; ok {
		p.Plus(move)
	}
}

func (d *Day2) Part2(moves [][]*util.Position2D) string {
	mp := make(map[util.Position2D]rune)
	mp[util.Position2D{X: -2, Y: 0}] = '5'
	mp[util.Position2D{X: -1, Y: 0}] = '6'
	mp[util.Position2D{X: 0, Y: 0}] = '7'
	mp[util.Position2D{X: 1, Y: 0}] = '8'
	mp[util.Position2D{X: 2, Y: 0}] = '9'
	mp[util.Position2D{X: 0, Y: -1}] = '3'
	mp[util.Position2D{X: -1, Y: -1}] = '2'
	mp[util.Position2D{X: 1, Y: -1}] = '4'
	mp[util.Position2D{X: 0, Y: -2}] = '1'
	mp[util.Position2D{X: 0, Y: 1}] = 'B'
	mp[util.Position2D{X: -1, Y: 1}] = 'A'
	mp[util.Position2D{X: 1, Y: 1}] = 'C'
	mp[util.Position2D{X: 0, Y: 2}] = 'D'
	str := []rune{}
	start := &util.Position2D{X: -2, Y: 0}
	for _, v := range moves {
		for _, q := range v {
			move(start, q, mp)
		}
		str = append(str, mp[*start])
	}

	return string(str)

}
