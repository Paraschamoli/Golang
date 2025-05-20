package main

import "fmt"

type Shape interface {
	Area() float32
}

type circle struct {
	radii float32
}
type rectangle struct {
	l float32
	b float32
}

func (c circle) Area() float32 {

	return 3.14 * (c.radii * c.radii)
}
func (r rectangle) Area() float32 {
	return r.b * r.l
}

func CalculateArea(s Shape)float32{
	return s.Area()
}

func main() {
	rec := rectangle{
		l: 5.5,
		b: 6.6,
	}
	cir:=circle{
		radii: 7.7,
	}
	fmt.Println("rectangle area =",CalculateArea(rec))
	fmt.Println("circle area= ",CalculateArea(cir))
}