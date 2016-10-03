package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var n int64 = 1
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the fibonnacci index to find the value:\n")
	input.Scan()
	x := input.Text()
	if number, err := strconv.ParseInt(x, 10, 0); err != nil {
		fmt.Println("\r%v is not a number.", x)
		os.Exit(0)
	} else {
		n = number
	}
	go spinner(100 * time.Millisecond)
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int64) int64 {
	if x < 2 {
		return x
	}
	val := fib(x-1) + fib(x-2)
	return val
}
