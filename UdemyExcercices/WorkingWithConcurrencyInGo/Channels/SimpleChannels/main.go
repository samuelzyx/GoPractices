package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping

		pong <- strings.ToUpper(s)
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type somenthing and press Enter (enter Q to quit)")

	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels")
	close(ping)
	close(pong)
}
