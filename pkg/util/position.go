package util

import "math"

type Position2D struct {
	X int
	Y int
}

func (p *Position2D) Distance() int {
	return Abs(p.X) + Abs(p.Y)
}

func (p *Position2D) Plus(other *Position2D) {
	p.X = p.X + other.X
	p.Y = p.Y + other.Y
}

func Sum(a, b *Position2D) *Position2D {
	return &Position2D{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func Diff(a, b *Position2D) *Position2D {
	return Sum(a, Scale(b, -1))
}

func Scale(a *Position2D, i int) *Position2D {
	return &Position2D{
		X: a.X * i,
		Y: a.Y * i,
	}
}

func GetDirection(a *Position2D) *Position2D {
	x := 0
	if a.X < 0 {
		x = -1
	} else if a.X > 0 {
		x = 1
	}
	y := 0
	if a.Y < 0 {
		y = -1
	} else if a.Y > 0 {
		y = 1
	}
	return &Position2D{
		X: x,
		Y: y,
	}
}

func (p *Position2D) Times(i int) {
	p.X = p.X * i
	p.Y = p.Y * i
}

func (p *Position2D) DistanceTo(other *Position2D) float64 {
	xdist := p.X - other.X
	ydist := p.Y - other.Y
	return math.Sqrt(float64(xdist*xdist) + float64(ydist*ydist))
}

var (
	Up    Position2D = Position2D{X: 0, Y: -1}
	Down  Position2D = Position2D{X: 0, Y: 1}
	Left  Position2D = Position2D{X: -1, Y: 0}
	Right Position2D = Position2D{X: 1, Y: 0}
)
