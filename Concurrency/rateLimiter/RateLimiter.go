package main

import (
	"fmt"
	"time"
)

func main() {

	//populate buffered channel with 5 entries.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request1", req, time.Now())
	}

	//type 2
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second * 1) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request2", req, time.Now())
	}
}
