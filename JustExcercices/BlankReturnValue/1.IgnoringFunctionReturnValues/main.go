package main

import "fmt"

// Exercise 1: Ignoring Function Return Values
// Imagine you have a function that calculates
// the area and perimeter of a rectangle, but
// you're only interested in the area. Here's the function:

func main() {
	area, _ := calculateAreaAndPerimeter(5, 5)
	fmt.Println("Area:", area)
}

func calculateAreaAndPerimeter(width float64, height float64) (float64, float64) {
	area := width * height
	perimeter := 2 * (width + height)
	return area, perimeter
}
