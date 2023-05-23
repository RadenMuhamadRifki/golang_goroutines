package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Helo")
	time.Sleep(1 * time.Second)
}

func TestAsync(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

var counter = 0

func OnceFunc() {
	counter++
}

func TestOnceData(t *testing.T) {
	once := &sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnceFunc)
			group.Done()
		}()

	}
	group.Wait()
	fmt.Println(counter)
}
