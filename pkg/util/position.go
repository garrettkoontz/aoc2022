package util

type Position2D struct {
	X int
	Y int
}

func (p *Position2D) Distance() int {
	return Abs(p.X) + Abs(p.Y)
}

func (p *Position2D) Plus(other Position2D) {
	p.X = p.X + other.X
	p.Y = p.Y + other.Y
}

func (p *Position2D) Times(i int) {
	p.X = p.X * i
	p.Y = p.Y * i
}
