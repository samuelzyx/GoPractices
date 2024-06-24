package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Println("Before swap function: ", a, b)
	swap(&a, &b) // Pass memory addresses of a and b using pointers (anpersan)
	fmt.Println("After swap function: ", a, b)
}

func swap(x *int, y *int) {
	temp := *x // Dereference x to get the value and store it in temp
	*x = *y    // Dereference y to get its value and assign it to the memory location pointed to by x
	*y = temp  // Assign the value originally stored in x (now in temp) to the memory location pointed to by y
}
