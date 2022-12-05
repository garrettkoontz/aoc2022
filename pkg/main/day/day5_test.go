package day

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d5 = Day5{}

var day5Input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func Test_d5part1(t *testing.T) {
	p := d.processInput(strings.Split(day5Input, "\n"))
	assert.Equal(t, "CMZ", d.part1(p), "equals CMZ")

}
