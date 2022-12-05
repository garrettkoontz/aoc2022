package day

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day4input = "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"

var d = Day4{}

func Test_part1(t *testing.T) {
	p := d.processInput(strings.Split(day4input, "\n"))
	assert.Equal(t, 2, d.part1(p), "equals 2")

}
