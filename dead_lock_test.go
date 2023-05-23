package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) UnLock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) ChangeAmount(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1 ", user1.Name)
	user1.ChangeAmount(-amount)
	time.Sleep(1 * time.Second)
	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.ChangeAmount(amount)
	time.Sleep(5 * time.Second)
	user1.UnLock()

	user2.UnLock()

}
func TestDeadLock(t *testing.T) {
	var group sync.WaitGroup
	defer group.Done()
	balance1 := UserBalance{
		Name:    "Rifki",
		Balance: 100000,
	}
	balance2 := UserBalance{
		Name:    "Raden",
		Balance: 90000,
	}

	go Transfer(&balance1, &balance2, 10000)
	go Transfer(&balance2, &balance1, 20000)

	time.Sleep(3 * time.Second)

	fmt.Println("User 1", balance1.Name, "Balance", balance1.Balance)
	fmt.Println("User 1", balance2.Name, "Balance", balance2.Balance)
	group.Wait()
}
