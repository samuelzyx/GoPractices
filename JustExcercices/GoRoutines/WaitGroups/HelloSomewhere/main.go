package main

import (
	"fmt"
	"sync"
)

var MSN string
var WG sync.WaitGroup

func updateMessage(s string) {
	defer WG.Done()
	MSN = s
}

func printMessage() {
	fmt.Println(MSN)
}

func main() {
	MSN = "Hello, world!"

	WG.Add(1)
	go updateMessage("Hello, universe!")
	WG.Wait()
	printMessage()

	WG.Add(1)
	go updateMessage("Hello, cosmos!")
	WG.Wait()
	printMessage()

	WG.Add(1)
	go updateMessage("Hello, world!")
	WG.Wait()
	printMessage()

}
