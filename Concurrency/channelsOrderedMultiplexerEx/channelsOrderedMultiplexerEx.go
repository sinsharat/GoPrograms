package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func first(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c //the value sent to c is either <-input1 or <-input2 depending upon who receives first
}

func gopher(msg string) <-chan Message { //receive channel of type Message
	c := make(chan Message)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-wait //waiting to receive a go ahead to reply
		}
	}()
	return c
}

func main() {
	c := first(gopher("Bat"), gopher("Robin"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true // go ahead bat
		msg2.wait <- true // go ahead robin
	}
	fmt.Println("You both talk too much. Bye!")
}
