package figures

type Square struct {
	Side float64
}

func (s *Square) getPerimeter() float64 {
	return s.Side * 4
}

func (s *Square) getArea() float64 {
	return s.Side * s.Side
}
