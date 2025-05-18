package banking3

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mutex   sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance += amount

	fmt.Printf("Deposited $%d. New Balance is $%d\n", amount, a.balance)

}

func (a *Account) Withdraw(amount int) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew $%d. New Balance is $%d\n", amount, a.balance)
		return true
	}

	fmt.Printf("Failed to withdraw $%d. Insufficient funds. Balance: $%d\n", amount, a.balance)
	return false
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.balance
}

func NewAccount(initialBalance int) *Account {
	return &Account{
		balance: initialBalance,
	}
}

func ManyDeposits(account *Account, amount int, numDeposits int) {

	var wg sync.WaitGroup

	for i := 0; i < numDeposits; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(amount)
		}()
	}
	wg.Wait()
}
