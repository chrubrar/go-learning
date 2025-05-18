package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/chrubrar/go-learning/banking"
)

func main() {
	// Create a new account with initial balance of 1000
	account := banking.NewAccount(1000)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch 5 deposit goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter
		go func(amount int) {
			defer wg.Done()                                      // Decrement counter when done
			time.Sleep(time.Millisecond * time.Duration(amount)) // Simulate processing time
			account.Deposit(amount * 100)
		}(i)
	}

	// Launch 3 withdrawal goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(amount*2)) // Simulate processing time
			account.Withdraw(amount * 300)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final balance
	fmt.Printf("\nFinal balance: %d\n", account.Balance())
}
