package main

import "fmt"

// Write a function that takes a boolean value and returns a
// string representation of that boolean ("true" or "false").
// You can use an if statement or a conditional expression for
// this conversion.
// In your main function, prompt the user to enter a value (e.g.,
// "yes" or "no"). Use string comparison operators to check
// if the input matches "yes" or "no". Convert the corresponding
// boolean value (true for "yes", false for "no") to a
// string using your function and print the result

func conversionBooleanToString(flag bool) string {
	if flag {
		return "yes"
	}
	return "no"
}

func main() {
	boolean1 := true
	boolean2 := false

	fmt.Println("If boolean var is", boolean1, "it means that it is", conversionBooleanToString(boolean1))
	fmt.Println("If boolean var is", boolean2, "it means that it is", conversionBooleanToString(boolean2))
}
