package aoc2022

import (
	"aoc2022/pkg/util"
	"fmt"
)

type Day12 struct {
}

type Map struct {
	grid  [][]rune
	start util.Position2D
	end   util.Position2D
}

func (m *Map) possibleSteps(p *util.Position2D, visited map[util.Position2D]bool) []*util.Position2D {
	result := []*util.Position2D{}
	for _, v := range []*util.Position2D{&util.Up, &util.Down, &util.Left, &util.Right} {
		w := util.Sum(p, v)
		if _, ok := visited[*w]; !ok &&
			w.X >= 0 &&
			w.Y >= 0 &&
			w.X < len(m.grid) &&
			w.Y < len(m.grid[0]) {
			r := m.grid[p.X][p.Y]
			if m.grid[w.X][w.Y] <= r+1 {
				result = append(result, w)
			}
		}
	}
	return result
}

type path struct {
	Path   map[util.Position2D]int
	Latest util.Position2D
}

func (p *path) next(ps []*util.Position2D) []*path {
	paths := []*path{}
	for _, v := range ps {
		pos := *v
		np := make(map[util.Position2D]int, len(p.Path))
		for k, v := range p.Path {
			np[k] = v
		}
		if _, ok := np[pos]; ok {
			continue
		}
		np[pos] = len(p.Path)
		paths = append(paths, &path{Path: np, Latest: pos})
	}
	return paths
}

func (m *Map) shortestPath() *path {
	queue := []*path{
		{
			Path: map[util.Position2D]int{m.start: 0}, Latest: m.start,
		},
	}
	visited := map[util.Position2D]bool{m.start: true}
	end := m.end
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		nexts := m.possibleSteps(&p.Latest, visited)
		newPaths := p.next(nexts)
		for _, v := range newPaths {
			visited[v.Latest] = true
			if v.Latest == end {
				return v
			}
		}
		queue = append(queue, newPaths...)
	}
	return nil
}

func (d Day12) processInput(input []string) *Map {
	result := make([][]rune, len(input))
	start := util.Position2D{X: 0, Y: 0}
	end := util.Position2D{X: 0, Y: 0}
	for i, line := range input {
		result[i] = make([]rune, len(line))
		for j, r := range line {
			result[i][j] = r
			if r == 'S' {
				result[i][j] = 'a'
				start = util.Position2D{X: i, Y: j}
			}
			if r == 'E' {
				result[i][j] = 'z'
				end = util.Position2D{X: i, Y: j}
			}
		}
	}
	return &Map{result, start, end}
}

func (d Day12) Run() {
	fileName := "input_2022/day12.txt"
	input, err := util.ReadLines(fileName)
	if err != nil {
		panic(err)
	}
	dir := d.processInput(input)
	fmt.Println(d.part1(dir))
	fmt.Println(d.part2(dir))
}

func (d Day12) part1(mp *Map) int {
	return len(mp.shortestPath().Path) - 1
}

func (d Day12) part2(mp *Map) int {
	lengths := []int{}
	for i := 0; i < len(mp.grid); i++ {
		for j := 0; j < len(mp.grid[i]); j++ {
			if mp.grid[i][j] == 'a' {
				newMp := &Map{
					grid:  mp.grid,
					start: util.Position2D{X: i, Y: j},
					end:   mp.end,
				}
				path := newMp.shortestPath()
				if path != nil {
					pathLength := len(path.Path) - 1
					lengths = append(lengths, pathLength)
				}
			}
		}
	}
	min := len(mp.grid) * len(mp.grid[0])
	for _, l := range lengths {
		if l < min {
			min = l
		}
	}
	return min
}
