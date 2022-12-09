package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"strconv"
)

type Day8 struct {
}

func (d Day8) processInput(input []string) [][]int {
	result := make([][]int, len(input))
	crossSize := len(input[0])
	for y, line := range input {
		sub := make([]int, crossSize)
		for x, c := range line {
			i, err := strconv.Atoi(string([]rune{c}))
			if err != nil {
				panic(err)
			}
			sub[x] = i
		}
		result[y] = sub
	}
	return result
}

func isVisible(grid [][]int, pos util.Position2D) bool {
	if pos.X == 0 || pos.Y == 0 || pos.X == len(grid[0]) || pos.Y == len(grid) {
		return true
	}
	value := grid[pos.Y][pos.X]
	isVisibleLeft := true
	for i := 0; i < pos.X; i++ {
		if grid[pos.Y][i] >= value {
			isVisibleLeft = false
			break
		}
	}
	isVisibleRight := true
	for i := pos.X + 1; i < len(grid[0]); i++ {
		if grid[pos.Y][i] >= value {
			isVisibleRight = false
			break
		}
	}
	isVisibleUp := true
	for i := 0; i < pos.Y; i++ {
		if grid[i][pos.X] >= value {
			isVisibleUp = false
			break
		}
	}
	isVisibleDown := true
	for i := pos.Y + 1; i < len(grid); i++ {
		if grid[i][pos.X] >= value {
			isVisibleDown = false
			break
		}
	}
	return isVisibleDown || isVisibleLeft || isVisibleRight || isVisibleUp
}

func scenicScore(grid [][]int, pos util.Position2D) int {
	if pos.X == 0 || pos.Y == 0 || pos.X == len(grid[0])-1 || pos.Y == len(grid)-1 {
		return 0
	}
	value := grid[pos.Y][pos.X]
	leftScore := pos.X
	for i := 1; i <= pos.X; i++ {
		if grid[pos.Y][pos.X-i] >= value {
			leftScore = i
			break
		}
	}
	rightScore := len(grid[0]) - pos.X - 1
	for i := 1; i < len(grid[0])-pos.X-1; i++ {
		if grid[pos.Y][pos.X+i] >= value {
			rightScore = i
			break
		}
	}
	upScore := pos.Y
	for i := 1; i <= pos.Y; i++ {
		if grid[pos.Y-i][pos.X] >= value {
			upScore = i
			break
		}
	}
	downScore := len(grid) - pos.Y - 1
	for i := 1; i < len(grid)-pos.Y-1; i++ {
		if grid[pos.Y+i][pos.X] >= value {
			downScore = i
			break
		}
	}
	return leftScore * rightScore * upScore * downScore
}

func (d Day8) Run() {
	fileName := "input_2022/day8.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	fmt.Println(d.part2(dir))
}

func (d Day8) part1(grid [][]int) int {
	countVisible := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if isVisible(grid, util.Position2D{X: x, Y: y}) {
				countVisible += 1
			}
		}
	}
	return countVisible
}

func (d Day8) part2(grid [][]int) int {
	score := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			newScore := scenicScore(grid, util.Position2D{X: x, Y: y})
			if newScore > score {
				score = newScore
			}
		}
	}
	return score
}

// too high 3683680
// too high 3439800
