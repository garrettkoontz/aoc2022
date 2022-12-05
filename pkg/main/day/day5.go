package day

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Day5 struct {
}

type StackInput struct {
	Stacks []stack
	Moves  []move
}

type move struct {
	Qty  int
	From int
	To   int
}

func (s *StackInput) makeMoves9000() {
	for _, m := range s.Moves {
		applyMove9000(s.Stacks, m)
	}
}

func applyMove9000(s []stack, m move) {
	fromStack := s[m.From]
	toStack := s[m.To]
	for i := 0; i < m.Qty; i++ {
		fs, r := fromStack.Pop()
		ts := toStack.Push(r)
		fromStack = fs
		toStack = ts
	}
	s[m.From] = fromStack
	s[m.To] = toStack
}

func (s *StackInput) makeMoves9001() {
	for _, m := range s.Moves {
		applyMove9001(s.Stacks, m)
	}
}

func applyMove9001(s []stack, m move) {
	fromStack := s[m.From]
	toStack := s[m.To]
	tempStack := stack{}
	for i := 0; i < m.Qty; i++ {
		fs, r := fromStack.Pop()
		ts := tempStack.Push(r)
		fromStack = fs
		tempStack = ts
	}
	for i := 0; i < m.Qty; i++ {
		tmps, r := tempStack.Pop()
		ts := toStack.Push(r)
		tempStack = tmps
		toStack = ts
	}
	s[m.From] = fromStack
	s[m.To] = toStack
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() rune {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[l-1]
}

func parseStacks(input []string) []stack {
	numStacks := (len(input[0]) + 1) / 4
	maxStackHeight := len(input) - 1
	stacks := make([]stack, numStacks)
	for i := maxStackHeight - 1; i >= 0; i-- {
		row := []rune(input[i])
		for j := 0; j < numStacks; j++ {
			r := row[(j*4)+1]
			if r != ' ' {
				stacks[j] = stacks[j].Push(r)
			}
		}
	}
	return stacks
}

func parseMoves(input []string) []move {
	result := make([]move, len(input))
	re := regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	for i, v := range input {
		groups := re.FindStringSubmatch(v)
		qty, err := strconv.Atoi(groups[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(groups[2])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(groups[3])
		if err != nil {
			panic(err)
		}
		result[i] = move{
			Qty:  qty,
			From: from - 1,
			To:   to - 1,
		}
	}
	return result
}

func (d Day5) processInput(input []string) *StackInput {
	bk := -1
	for i, v := range input {
		if v == "" {
			bk = i
		}
	}
	stacksInput := input[:bk]
	movesInput := input[bk+1:]

	return &StackInput{
		Stacks: parseStacks(stacksInput),
		Moves:  parseMoves(movesInput),
	}
}

func (d Day5) Run() {
	start := time.Now()
	// Code to measure
	fileName := "input/Day5.txt"
	input, err := ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	inp := d.processInput(input)
	start1 := time.Now()
	fmt.Println(d.part1(inp))
	duration := time.Since(start1)
	fmt.Println(duration)
	inp = d.processInput(input)
	start1 = time.Now()
	fmt.Println(d.part2(inp))
	duration = time.Since(start1)
	fmt.Println(duration)
	duration = time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
}

func (d Day5) part1(i *StackInput) string {
	i.makeMoves9000()
	s := []rune{}
	for _, is := range i.Stacks {
		s = append(s, is.Peek())
	}
	return string(s)
}

func (d Day5) part2(i *StackInput) string {
	i.makeMoves9001()
	s := []rune{}
	for _, is := range i.Stacks {
		s = append(s, is.Peek())
	}
	return string(s)
}
