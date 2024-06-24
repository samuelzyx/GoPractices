package main

import "fmt"

//Write a function that takes an interface{} and attempts to
// type assert it to a specific type (e.g., string). Use a switch
// statement with type assertions to handle different potential
// types for the interface value. Print a message indicating the
// success or failure of the type assertion and the actual type of the value.

func attemptConversion(dataType interface{}) {
	switch typeOf := dataType.(type) {
	case string:
		fmt.Println("This is a string:", typeOf)
	case bool:
		fmt.Println("This is a bool:", typeOf)
	case int:
		fmt.Println("This is a int:", typeOf)
	case float32:
		fmt.Println("This is a float32:", typeOf)
	case float64:
		fmt.Println("This is a float64:", typeOf)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	attemptConversion(1)
	attemptConversion(1.25)
	attemptConversion("Sam")
	attemptConversion(false)
}
