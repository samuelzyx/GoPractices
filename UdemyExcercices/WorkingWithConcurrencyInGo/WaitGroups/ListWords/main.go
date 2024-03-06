package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printLetter(fmt.Sprintf("%d %s", i, x), &wg)
	}

	wg.Wait()
}

func printLetter(l string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(l)
}
