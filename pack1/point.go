package pack1

import "math"

type Point struct {
	x, y float64
}

func (p *Point) SetXY(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	Name string
}
