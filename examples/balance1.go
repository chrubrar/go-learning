package examples

import (
	"fmt"
	"sync"
	"time"
)

// BankAccount WITHOUT mutex (unsafe)
type UnsafeBankAccount struct {
	balance int
}

func (b *UnsafeBankAccount) Deposit(amount int) {
	// Simulate some processing time
	//time.Sleep(time.Millisecond)
	b.balance = b.balance + amount
}

// BankAccount WITH mutex (safe)
type SafeBankAccount struct {
	balance int
	mutex   sync.Mutex // This is the mutex - think of it as a lock
}

func (b *SafeBankAccount) Deposit(amount int) {
	b.mutex.Lock()         // Lock the door - only one goroutine can enter
	defer b.mutex.Unlock() // Unlock when function returns

	// Simulate some processing time
	time.Sleep(time.Millisecond)
	b.balance = b.balance + amount
}

func main() {
	// Example 1: Without Mutex (UNSAFE)
	unsafeAccount := &UnsafeBankAccount{balance: 0}

	// Start 100 concurrent deposits of $100 each
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeAccount.Deposit(100)
		}()
	}
	wg.Wait()

	fmt.Printf("Unsafe Account Balance (should be 10000): %d\n", unsafeAccount.balance)

	// Example 2: With Mutex (SAFE)
	safeAccount := &SafeBankAccount{balance: 0}

	// Start 100 concurrent deposits of $100 each
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeAccount.Deposit(100)
		}()
	}
	wg.Wait()

	fmt.Printf("Safe Account Balance (will be 10000): %d\n", safeAccount.balance)
}
