package banking

import (
	"fmt"
	"sync"
)

// Account represents a bank account with thread-safe operations
type Account struct {
	balance int
	mutex   sync.Mutex // protects balance from concurrent access
}

// NewAccount creates a new account with initial balance
func NewAccount(initialBalance int) *Account {
	return &Account{
		balance: initialBalance,
	}
}

// Deposit adds money to the account
func (a *Account) Deposit(amount int) {
	a.mutex.Lock()         // Lock before accessing balance
	defer a.mutex.Unlock() // Ensure unlock happens after function completes

	a.balance += amount
	fmt.Printf("Deposited %d. New balance: %d\n", amount, a.balance)
}

// Withdraw removes money from the account if sufficient funds exist
func (a *Account) Withdraw(amount int) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d. New balance: %d\n", amount, a.balance)
		return true
	}

	fmt.Printf("Failed to withdraw %d. Insufficient funds. Balance: %d\n", amount, a.balance)
	return false
}

// Balance returns the current balance
func (a *Account) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.balance
}
