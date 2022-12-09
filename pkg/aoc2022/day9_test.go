package aoc2022

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var d9 = Day9{}

func Test_d9part1(t *testing.T) {
	inp := d9.processInput(strings.Split(input, "\n"))
	assert.Equal(t, 13, d9.part1(inp), "Equals 13")
}

func Test_d9part2(t *testing.T) {
	inp := d9.processInput(strings.Split(input, "\n"))
	assert.Equal(t, 1, d9.part2(inp), "Equals 1")
}
