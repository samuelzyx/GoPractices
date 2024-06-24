package main

import "fmt"

type Shape interface {
	ShapeName() string
	Area() float64
	Perimeter() float64
}

type Square struct {
	Name string
	Side float64
}

type Circle struct {
	Name   string
	Radius float64
}

func (s Square) ShapeName() string {
	return s.Name
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func (c Circle) ShapeName() string {
	return c.Name
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func Calculate(s Shape) {
	fmt.Println("Area for", s.ShapeName(), ":", s.Area())
	fmt.Println("Perimeter for", s.ShapeName(), ":", s.Perimeter())
}

func main() {
	square := &Square{Name: "Square", Side: 5}
	circle := &Circle{Name: "Circle", Radius: 10}

	fmt.Println("Square, with side:", square.Side)
	fmt.Println("Circle, with radius:", circle.Radius)

	Calculate(square)
	Calculate(circle)
}
