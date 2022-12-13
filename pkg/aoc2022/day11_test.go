package aoc2022

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m0 = Monkey{0, []uint64{79, 98}, Operation{"*", 19}, 23, 2, 3, 0}
var m1 = Monkey{1, []uint64{54, 65, 75, 74}, Operation{"+", 6}, 19, 2, 0, 0}
var m2 = Monkey{2, []uint64{79, 60, 97}, Operation{"2", 2}, 13, 1, 3, 0}
var m3 = Monkey{3, []uint64{74}, Operation{"+", 3}, 17, 0, 1, 0}

var m = map[int]*Monkey{
	0: &m0,
	1: &m1,
	2: &m2,
	3: &m3,
}

var d11 = Day11{}

func Test_d11part1(t *testing.T) {
	assert.Equal(t, 10605, d11.part1(m), "equals")
}

func Test_d11part2(t *testing.T) {
	md := uint64(1)
	for _, mon := range m {
		md *= mon.Test
	}
	for i := 0; i < 20; i++ {
		round(m, func(u uint64) uint64 { return u % md })
	}
	for _, m := range m {
		fmt.Printf("Monkey: %d, times: %d\n", m.Number, m.inspectedCount)
	}
	assert.Equal(t, 2713310158, d11.part2(m), "equals")
}
