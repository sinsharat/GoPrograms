package main

import (
	"fmt"
	"time"
)

func main() {

	// Buffered channel blocks new writes only when the buffer is full.
	// Reciever receives block when the buffered channel is empty.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
}
