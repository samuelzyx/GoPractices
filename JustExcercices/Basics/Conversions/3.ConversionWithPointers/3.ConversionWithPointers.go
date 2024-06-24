package main

import "fmt"

// Write a function that takes an integer pointer (*int)
// and increments the value pointed to by that pointer.
// In your main function, declare an integer variable and
// initialize it with a value. Pass the address of this
// variable (using the & operator) to your function. After
// calling the function, print the original variable's
// value again to see the modification.

func incrementsIntValueByPointer(counter *int) { //affect value from memory address
	*counter += 10
}

func main() {
	counter := 34
	fmt.Println("Value of counter", counter)
	incrementsIntValueByPointer(&counter) //sending memory address
	fmt.Println("Incremented to", counter)
}
