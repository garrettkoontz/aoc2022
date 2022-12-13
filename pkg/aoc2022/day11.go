package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day11 struct {
}

type Monkey struct {
	Number         int
	Items          []uint64
	Operation      Operation
	Test           uint64
	TrueThrow      int
	FalseThrow     int
	inspectedCount uint64
}

type Operation struct {
	Op    string
	Value uint64
}

func round(monkeys map[int]*Monkey, worryModulator func(uint64) uint64) {
	for j := 0; j < len(monkeys); j++ {
		monkeys[j].inspect(monkeys, worryModulator)
	}
}

func (m *Monkey) inspect(others map[int]*Monkey, worryModulator func(uint64) uint64) {
	for _, v := range m.Items {
		t := v
		m.inspectedCount++
		if m.Operation.Op == "+" {
			t += m.Operation.Value
		} else if m.Operation.Op == "*" {
			t *= m.Operation.Value
		} else if m.Operation.Op == "2" {
			t *= t
		} else {
			panic("Unknown operation")
		}
		t = worryModulator(t)
		if t%m.Test == 0 {
			others[m.TrueThrow].Items = append(others[m.TrueThrow].Items, t)
		} else {
			others[m.FalseThrow].Items = append(others[m.FalseThrow].Items, t)
		}
	}
	m.Items = []uint64{}
}

func (d Day11) processInput(input string) map[int]*Monkey {
	result := map[int]*Monkey{}
	spl := regexp.MustCompile(`\n\n`).Split(input, -1)

	digitRegex := regexp.MustCompile(`\d+`)
	for _, m := range spl {
		sp := strings.Split(m, "\n")
		num := -1
		items := []uint64{}
		operation := Operation{"", 0}
		test := -1
		tr := -1
		fl := -1
		var err error
		for i, l := range sp {
			switch i {
			case 0:
				num, err = strconv.Atoi(digitRegex.FindString(l))
				if err != nil {
					panic(err)
				}
			case 1:
				{
					itemsParse := digitRegex.FindAllString(l, -1)
					items = make([]uint64, len(itemsParse))
					for i, item := range itemsParse {
						it, err := strconv.Atoi(item)
						if err != nil {
							panic(err)
						}
						items[i] = uint64(it)
					}
				}
			case 2:
				operand := digitRegex.FindString(l)
				op := "2"
				o := 2
				if operand != "" {
					op = regexp.MustCompile(`[*+]`).FindString(l)
					o, err = strconv.Atoi(operand)
				}
				if err != nil {
					panic(err)
				}
				operation = Operation{op, uint64(o)}
			case 3:
				test, err = strconv.Atoi(digitRegex.FindString(l))
			case 4:
				tr, err = strconv.Atoi(digitRegex.FindString(l))
			case 5:
				fl, err = strconv.Atoi(digitRegex.FindString(l))
			}
		}
		result[num] = &Monkey{
			num, items, operation, uint64(test), tr, fl, 0,
		}
	}
	return result
}

func (d Day11) Run() {
	fileName := "input_2022/day11.txt"
	input, err := util.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	dir = d.processInput(input)
	fmt.Println(d.part2(dir))
}

func (d Day11) part1(monkeys map[int]*Monkey) uint64 {
	for _, m := range monkeys {
		fmt.Printf("Monkey: %+v\n", m)
	}
	for i := 0; i < 20; i++ {
		round(monkeys, func(u uint64) uint64 { return u / 3 })
	}
	inspCount := make([]uint64, len(monkeys))
	for i, m := range monkeys {
		inspCount[i] = m.inspectedCount
	}
	sort.Slice(inspCount, func(i, j int) bool {
		return inspCount[i] > inspCount[j]
	})
	fmt.Println(inspCount)
	return inspCount[0] * inspCount[1]
}

func (d Day11) part2(monkeys map[int]*Monkey) uint64 {
	m := uint64(1)
	for _, mon := range monkeys {
		m *= mon.Test
	}
	for i := 0; i < 10000; i++ {
		round(monkeys, func(u uint64) uint64 { return u % m })
	}
	inspCount := make([]uint64, len(monkeys))
	for i, m := range monkeys {
		inspCount[i] = m.inspectedCount
	}
	sort.Slice(inspCount, func(i, j int) bool {
		return inspCount[i] > inspCount[j]
	})
	fmt.Println(inspCount)
	return inspCount[0] * inspCount[1]
}

// too high 512640

// part 2 too high 18394603950

// part 2 too high 18297957888
