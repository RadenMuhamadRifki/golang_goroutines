package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	data := sync.Map{}
	var AddToMap = func(value int) {
		data.Store(value, value)
	}
	for i := 0; i < 100; i++ {
		go AddToMap(i)
	}
	time.Sleep(3 * time.Second)

	data.Range(func(key, value interface{}) bool {
		fmt.Println("key :", key, "value :", value)
		return true
	})
}
