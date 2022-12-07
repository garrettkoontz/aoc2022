package aoc2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input1 = "vJrwpWtwJgWrhcsFMMfFFhFp"
const input2 = "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
const input3 = "PmmdzqPrVvPwwTWBwg"
const input4 = "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"
const input5 = "ttgJtRGJQctTZtZT"
const input6 = "CrZsJsPPZsGzwwsLwLmpwMDw"

func Test_share(t *testing.T) {
	d := Day3{}

	inp1 := d.processInput([]string{input1})
	assert.Equal(t, 'p', inp1[0].share(), "is p")

	inp2 := d.processInput([]string{input2})
	assert.Equal(t, 'L', inp2[0].share(), "is p")

}

func Test_priority(t *testing.T) {
	assert.Equal(t, 16, priority('p'))
	assert.Equal(t, 38, priority('L'))
}
