package main

import (
	"fmt"
	"strings"
)

// Write a function that takes a string containing comma-separated
// values and converts it into a slice of strings. Use the strings.
// Split function to achieve this.
// In your main function, prompt the user to enter a comma-separated
// list of words (e.g., "apple,banana,orange"). Call your function
// to convert the input string into a slice of strings and print
// the elements of the slice.

func main() {
	entry := "apple,banana,orange"
	slice := strings.Split(entry, ",")

	for _, value := range slice {
		fmt.Println(value)
	}
}
