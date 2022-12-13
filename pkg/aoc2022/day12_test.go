package aoc2022

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d12Input = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

var d12 = Day12{}

func Test_d12Part1(t *testing.T) {
	a := d12.part1(d12.processInput(strings.Split(d12Input, "\n")))
	assert.Equal(t, 31, a, "something")
}
