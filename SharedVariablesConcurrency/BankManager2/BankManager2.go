// This Package provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// balance is confined to teller goroutine
var balance int
var mu sync.RWMutex

func Deposit(amount int) int {
	mu.RLock()
	defer mu.RUnlock()
	return deposit(amount)
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) (int, bool) {
	mu.RLock()
	defer mu.RUnlock()
	balance := deposit(-amount)
	if balance < 0 {
		balance = deposit(amount)
		return balance, false
	}
	return balance, true
}

func deposit(amount int) int {
	balance += amount
	return balance
}

func main() {
	for i := 1; i <= 100; i++ {
		go func() {
			depositAmt := i * 100
			fmt.Fprintf(os.Stdout, "Amount %v deposited successfully and balance in account is : %v \n", depositAmt, Deposit(depositAmt))
		}()

		go func() {
			bal := Balance()
			fmt.Println("Current balance in the account is : ", bal)
		}()

		go func() {
			r := rand.Intn(10) + 1
			withdrawAmt := i * r * 100
			balance, isSuccessful := Withdraw(withdrawAmt)
			if isSuccessful {
				fmt.Fprintf(os.Stdout, "Amount %v withdrawn successfully and current balance in account is : %v \n", withdrawAmt, balance)
			} else {
				fmt.Fprintf(os.Stdout, "Failed to withdraw amount %v as the current balance in account is only : %v \n", withdrawAmt, balance)
			}
		}()
	}
}
