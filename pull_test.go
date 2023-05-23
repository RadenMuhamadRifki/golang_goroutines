package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Rifki"
		},
	}
	group := sync.WaitGroup{}
	defer group.Done()
	pool.Put("Raden")
	pool.Put("Muhamad")
	pool.Put("Rifki")
	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	group.Wait()
	fmt.Println("Selesai")
}
