package main

import (
	"fmt"
	"time"
)

var counter int

func increment(ch chan int) {
	ch <- 1 // Send 1 to the channel
}

func main() {
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		go increment(ch)
	}

	go func() {
		for i := 0; i < 100; i++ {
			counter += <-ch // Receive from the channel and increment counter
		}
	}()

	// Wait for all goroutines to finish (including the listening one)
	time.Sleep(time.Millisecond)

	fmt.Println("Final counter (with channel):", counter)
}
