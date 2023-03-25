package figures

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) getArea() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
