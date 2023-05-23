package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()

		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)
	group.Wait()
}
func TestAddThread(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()

		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)
	group.Wait()

}
