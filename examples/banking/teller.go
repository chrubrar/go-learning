package main

import (
	"fmt"
	"github.com/chrubrar/go-learning/banking3"
	"sync"
	"time"
)

func main() {
	// Create a new account with initial balance of 1000
	account := banking3.NewAccount(0)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(i))
			account.Deposit(100)
		}()
	}
	wg.Wait()

	var wg1 sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			time.Sleep(time.Millisecond * time.Duration(i))
			account.Withdraw(300)
		}()
	}

	wg1.Wait()

	fmt.Printf("Final balance: $%d\n", account.Balance())

	// Make 100 deposits of 100 each
}
