package aoc2022

import (
	"aoc2022/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d8Input [][]int = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

var d8 = Day8{}

func Test_part2(t *testing.T) {
	assert.Equal(t, 8, d8.part2(d8Input), "equals 8")
	assert.Equal(t, 8, scenicScore(d8Input, util.Position2D{X: 2, Y: 3}))
	assert.Equal(t, 4, scenicScore(d8Input, util.Position2D{X: 2, Y: 1}))
}
