package main

import (
	"fmt"
	"sync"
)

var msn string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msn = s
	m.Unlock()
}

func main() {

	msn = "Hello, world!"

	var mutext sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutext)
	go updateMessage("Hello, cosmos!", &mutext)

	wg.Wait()

	fmt.Println(msn)
}
