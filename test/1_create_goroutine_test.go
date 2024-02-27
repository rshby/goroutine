package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// function yang akan dijalankan dengan goroutine
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// panggil function menggunakan goroutine
	go RunHelloWorld()
	fmt.Println("ups!")
	time.Sleep(1 * time.Second)
}

// buat function yang akan dirun menggunakan goroutine
func DisplayNumber(number int, mtx *sync.Mutex) {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println("display", number)
}

func TestManyGoroutine(t *testing.T) {
	mtx := &sync.Mutex{}

	// jalankan goroutine sebanyak 100000 thread
	for i := 1; i < 100000; i++ {
		go DisplayNumber(i, mtx)
	}

	time.Sleep(4 * time.Second)
}
