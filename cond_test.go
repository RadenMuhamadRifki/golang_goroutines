package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCond(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCond(i)
	}
	go func() {
		for j := 0; j < 10; j++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()

	}()
	group.Wait()
}

func TestTimer(t *testing.T) {
	chanel := time.After(3 * time.Second)
	fmt.Println(time.Now())
	time := <-chanel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Function Ini Akan Berjalan ketika 3 detik")
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}
func TestTicker(t *testing.T) {
	ticker := time.Tick(1 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
	}()
	for tick := range ticker {
		fmt.Println(tick)
	}

}
