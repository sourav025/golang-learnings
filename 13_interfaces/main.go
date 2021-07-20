package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 4}
	rec := Rectangle{4, 5}
	fmt.Printf("Area of circle = %.2f\n", getArea(circle))
	fmt.Printf("Area of Rectangle = %.2f\n", getArea(rec))

}
