package aoc2016_test

import (
	"aoc2022/pkg/aoc2016"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d1 aoc2016.Day1 = aoc2016.Day1{}

func Test_part1(t *testing.T) {
	r2 := aoc2016.Move{Direction: 'R', Amount: 2}
	l3 := aoc2016.Move{Direction: 'L', Amount: 3}
	assert.Equal(t, 5, d1.Part1([]*aoc2016.Move{
		&r2,
		&l3,
	}), "1")
}

func Test_part2(t *testing.T) {
	r8 := aoc2016.Move{Direction: 'R', Amount: 8}
	r4 := aoc2016.Move{Direction: 'R', Amount: 4}
	assert.Equal(t, 4, d1.Part2([]*aoc2016.Move{
		&r8,
		&r4,
		&r4,
		&r8,
	}), "1")
}
