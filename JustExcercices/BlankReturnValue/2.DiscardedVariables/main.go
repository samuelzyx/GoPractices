package main

import "fmt"

// Let's say you're looping through a list of
// names and only need to print every other name.
// How can you utilize the blank identifier to
// achieve this while still iterating through the entire list?

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Diana"}

	counter := 0

	for _, value := range names {
		if counter%2 == 0 {
			fmt.Println("Name:", value, "with index", counter) // Print only when count is even (every other name)
		}
		counter++
	}
}
