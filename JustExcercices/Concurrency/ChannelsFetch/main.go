package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2) //string and int match
	wg := &sync.WaitGroup{}     //wait group to sync go routines

	wg.Add(2)                               //set number of go routines to wait
	go fetchUserLikes(userName, respch, wg) //set this line as go routine
	go fetchUserMatch(userName, respch, wg) //set this line as go routine

	wg.Wait()     //force to sync and wait until each number go routines are performance
	close(respch) //close chan

	for resp_ := range respch {
		fmt.Println("resp", resp_) //print values received by chan
	}

	fmt.Println("took: ", time.Since(start)) //time taked to perform all go routines
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 1000)
	return "BOB"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	fmt.Println("Likes from user:", userName)
	time.Sleep(time.Millisecond * 1000) //time taked to perform this func
	respch <- 11                        //return value throw chan
	wg.Done()                           //notice wait group action done
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	fmt.Println("User match from user:", userName)
	time.Sleep(time.Millisecond * 1000) //time taked to perform this func
	respch <- "ANNA"                    //return value throw chan
	wg.Done()                           //notice wait group action done
}
