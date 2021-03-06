package main

import (
	"fmt"
)

func first(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
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

func gopher(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func main() {
	c := first(gopher("Bat"), gopher("Robin"))
	for i := 0; i < 10; i++ {
		fmt.Printf("I am Gopher%v\n", <-c)
	}
	fmt.Println("You both talk too much. Bye!")
}
