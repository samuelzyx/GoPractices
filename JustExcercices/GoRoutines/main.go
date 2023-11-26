package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("one")
	c := make(chan string) //create channel to use communication response of go routine
	go testFunction(c)     //run in background independly the func main
	areWeFinish := <-c     //we depend until response is assinged if not return nothing it will deadlock
	fmt.Printf("already finish: %v\n", areWeFinish)
	fmt.Println("two")
}

func testFunction(c chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Checking...")
		time.Sleep(1 * time.Second)
	}
	c <- "Over" //response communication of go routine func
}
