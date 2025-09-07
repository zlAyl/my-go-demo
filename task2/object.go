package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (r *Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}
func (r *Circle) Perimeter() float64 {
	return 2 * math.Pi * r.Radius
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle得周长是:", rectangle.Perimeter())
	fmt.Println("Rectangle得面积是:", rectangle.Area())

	circle := Circle{Radius: 5}

	fmt.Println("circle周长是:", circle.Perimeter())
	fmt.Println("circle得面积是:", circle.Area())

}
