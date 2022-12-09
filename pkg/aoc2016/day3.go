package aoc2016

import (
	"aoc2022/pkg/util"
	"fmt"
	"regexp"
	"strconv"
)

type Day3 struct {
}

type tri struct {
	a int
	b int
	c int
}

func (t *tri) possible() bool {
	// fmt.Printf("tri %+v", t)
	if t.a+t.b > t.c && t.b+t.c > t.a && t.c+t.a > t.b {
		return true
	}
	return false
}

func (d *Day3) processInput(input []string) []*tri {
	result := make([]*tri, len(input))
	for i, v := range input {
		regexp := regexp.MustCompile(`[\s]*([0-9]+)[\s]*([0-9]+)[\s]*([0-9]+)[\s]*`)
		s := regexp.FindStringSubmatch(v)
		a, err := strconv.Atoi(s[1])
		if err != nil {
			err = fmt.Errorf("Unable to convert %q; err: %w", v, err)
			panic(err)
		}
		b, err := strconv.Atoi(s[2])
		if err != nil {
			err = fmt.Errorf("Unable to convert %q; err: %w", v, err)
			panic(err)
		}
		c, err := strconv.Atoi(s[3])
		if err != nil {
			err = fmt.Errorf("Unable to convert %q; err: %w", v, err)
			panic(err)
		}
		t := tri{a, b, c}
		result[i] = &t
	}
	return result
}

func (d *Day3) Run() {
	fileName := "input_2016/day3.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)

	fmt.Println(d.Part1(inp))
	fmt.Println(d.Part2(inp))
}

func (d *Day3) Part1(inp []*tri) int {
	valid := 0
	for _, v := range inp {
		if v.possible() {
			valid++
		}
	}
	return valid
}

func (d *Day3) Part2(inp []*tri) int {
	newTri := make([]*tri, len(inp))
	for i := 0; i < len(inp); i += 3 {
		t1 := inp[i]
		t2 := inp[i+1]
		t3 := inp[i+2]
		nt1 := tri{t1.a, t2.a, t3.a}
		nt2 := tri{t1.b, t2.b, t3.b}
		nt3 := tri{t1.c, t2.c, t3.c}
		newTri[i] = &nt1
		newTri[i+1] = &nt2
		newTri[i+2] = &nt3
	}
	return d.Part1(newTri)
}
