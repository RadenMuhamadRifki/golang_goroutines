package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func sayHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(num int) {
	fmt.Println("Number ", num)
}
func TestCreateGoroutine(t *testing.T) {
	go sayHelloWorld()
	fmt.Println("uPSS")
	time.Sleep(1 * time.Second)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	fmt.Println("Ini Di bawah")
	time.Sleep(10 * time.Second)
}
func TestCreateChanel(t *testing.T) {

	chanel := make(chan string)
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "Raden Muhamad Rifki"
		fmt.Println("Selesai Mengirim Data Ke Channel")
	}()
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func giveMeResponse(chanel chan<- string) {
	time.Sleep(2 * time.Second)
	chanel <- "Raden Muhamad Rifki"
}
func getMeResponse(chanel <-chan string) {
	data := <-chanel
	fmt.Println(data)
}
func TestResponse(t *testing.T) {
	chanel := make(chan string)
	go giveMeResponse(chanel)
	go getMeResponse(chanel)
	time.Sleep(5 * time.Second)
	defer close(chanel)

}
func TestBufferdChanel(t *testing.T) {
	chanel := make(chan string, 3)
	defer close(chanel)
	chanel <- "Rifki"
	time.Sleep(5 * time.Second)

}

func TestRangeChanel(t *testing.T) {
	chanel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "Mengirim Data Ke " + "" + strconv.Itoa(i)
		}
		close(chanel)

	}()
	for data := range chanel {
		fmt.Println("Menerima Data", data)
	}

}

func TestSelectChanel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)
	go giveMeResponse(chanel1)
	go getMeResponse(chanel2)
	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data ke Chanel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data ke Chanel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}

}

func TestRaceChanel(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1

			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
