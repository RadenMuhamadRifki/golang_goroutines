package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	sync.RWMutex
	Balance int
}

func (account *BankAccount) ChangeBankAccountMount(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) ReadBankAaccountmount() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}
func TestReadWriteMutex(t *testing.T) {
	balance := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				balance.ChangeBankAccountMount(1)
				fmt.Println(balance.ReadBankAaccountmount())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance", balance.Balance)
}
