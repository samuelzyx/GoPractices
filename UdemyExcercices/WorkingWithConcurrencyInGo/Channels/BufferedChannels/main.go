package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		//simulate do a lot of stuffs
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10) //10 the mount of goroutines working at same time

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done")
	close(ch)
}
