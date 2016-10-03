package main

import (
	"fmt"
)

func main() {
	id := make(chan string)
	counter := 0
	go func() {
		for {
			id <- fmt.Sprintf("%x", counter)
			counter++
		}
	}()

	fmt.Printf("Id1 : %v\n", <-id)
	fmt.Printf("Id2 : %v\n", <-id)
}
