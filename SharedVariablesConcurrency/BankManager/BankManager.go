// This Package provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"math/rand"
	"os"
)

var deposits = make(chan int)   // send amount to deposit
var withdrawls = make(chan int) // recieve amount on withdrawal
var balances = make(chan int)   // receive balance
var result = make(chan string)  // receive balance

func Deposit(amount int) string {
	deposits <- amount
	return <-result
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) string {
	withdrawls <- amount
	return <-result
}

func teller() {
	// balance is confined to teller goroutine
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Fprintln(os.Stdin, "Came Here")
			result <- fmt.Sprintf("Deposited the amount %v successfully. Balance is : %v", amount, balance)
		case amount := <-withdrawls:
			fmt.Fprintln(os.Stdin, "Came Here")
			if balance >= amount {
				balance -= amount
				result <- fmt.Sprintf("Withdrew the amount %v successfully. Balance is : %v", amount, balance)
			} else {
				result <- fmt.Sprintf("Cannot withdraw the amount : %v. Balance is only : %v", amount, balance)
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	for i := 1; i <= 100; i++ {
		go func() {
			depositAmt := i * 100
			fmt.Fprintln(os.Stdout, Deposit(depositAmt))
		}()

		go func() {
			bal := Balance()
			fmt.Println("Current balance is : ", bal)
		}()

		go func() {
			r := rand.Intn(10) + 1
			withdrawAmt := i * r * 100
			fmt.Fprintln(os.Stdout, Withdraw(withdrawAmt))
		}()
	}
}
