package main

import (
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(20, 20)
	i := 0
	start := time.Now()
	for time.Since(start).Seconds() < 30 {
		if limiter.Allow() {
			log.Printf("Printing : %v\n", i)
			i++
		}
	}
}
