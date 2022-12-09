package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day9 struct {
}

type ropeMove struct {
	Direction *util.Position2D
	Qty       int
}

type rope1 struct {
	Head *util.Position2D
	Tail *util.Position2D
}

type rope2 struct {
	Knots []*util.Position2D
}

func sameVertical(a, b *util.Position2D) bool {
	return a.X == b.X && a.Y != b.Y
}

func sameHorizontal(a, b *util.Position2D) bool {
	return a.X != b.X && a.Y == b.Y
}

func resolveTail(a, b *util.Position2D) *util.Position2D {
	if sameVertical(a, b) || sameHorizontal(a, b) {
		if a.DistanceTo(b) > float64(1) {
			b.Plus(util.GetDirection(util.Sum(a, util.Scale(b, -1))))
			return b
		}
	}
	if a.DistanceTo(b) > math.Sqrt(2) {
		m := util.GetDirection(util.Diff(a, b))
		b.Plus(m)
		return b
	}
	return nil
}

func (r *rope1) applyRopeMove(rm *ropeMove) []util.Position2D {
	result := []util.Position2D{}
	for i := 0; i < rm.Qty; i++ {
		r.Head.Plus(rm.Direction)
		rt := resolveTail(r.Head, r.Tail)
		if rt != nil {
			result = append(result, *rt)
		}
	}
	return result
}

func (r *rope2) applyRopeMove(rm *ropeMove) []util.Position2D {
	result := []util.Position2D{}
	for i := 0; i < rm.Qty; i++ {
		r.Knots[0].Plus(rm.Direction)
		for j := 1; j < len(r.Knots); j++ {
			resolveTail(r.Knots[j-1], r.Knots[j])
		}
		result = append(result, *r.Knots[9])
	}
	return result
}

func (d Day9) processInput(input []string) []*ropeMove {
	result := make([]*ropeMove, len(input))
	for i, l := range input {
		s := strings.Split(l, " ")
		d := util.Up
		switch s[0] {
		case "L":
			d = util.Left
			break
		case "R":
			d = util.Right
			break
		case "D":
			d = util.Down
			break
		}
		amt, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		result[i] = &ropeMove{
			&d, amt,
		}
	}
	return result
}

func (d Day9) Run() {
	fileName := "input_2022/day9.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	fmt.Println(d.part2(dir))
}

func (d Day9) part1(moves []*ropeMove) int {
	mp := make(map[util.Position2D]bool)
	mp[util.Position2D{X: 0, Y: 0}] = true
	rope := &rope1{
		Head: &util.Position2D{X: 0, Y: 0},
		Tail: &util.Position2D{X: 0, Y: 0},
	}
	for _, move := range moves {
		newTails := rope.applyRopeMove(move)
		for _, nt := range newTails {
			mp[nt] = true
		}
	}
	return len(mp)
}

func (d Day9) part2(moves []*ropeMove) int {
	mp := make(map[util.Position2D]bool)
	mp[util.Position2D{X: 0, Y: 0}] = true
	rope := &rope2{
		Knots: []*util.Position2D{
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 0},
		},
	}
	for _, move := range moves {
		newTails := rope.applyRopeMove(move)
		for _, nt := range newTails {
			mp[nt] = true
		}
	}
	return len(mp)
}

// too high 11220
