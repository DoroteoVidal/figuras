package figures

import "fmt"

type DataFigure interface {
	getArea() float64
	getPerimeter() float64
}

func Measures(d DataFigure) {
	fmt.Println(d)
	fmt.Println(d.getArea())
	fmt.Println(d.getPerimeter())
}
