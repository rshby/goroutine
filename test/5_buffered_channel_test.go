package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	// create channel with buffer
	chanString := make(chan string, 100)

	fmt.Println("cap channel :", cap(chanString))
	fmt.Println("panjang channel :", len(chanString))
	fmt.Println()

	// go func
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(chanString chan<- string) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			chanString <- fmt.Sprintf("%v", i+1)
		}

		close(chanString)
	}(chanString)

	wg.Add(1)
	go func(wg *sync.WaitGroup, chanString chan string) {
		defer wg.Done()
		for data := range chanString {
			fmt.Println(data)
		}
	}(wg, chanString)

	wg.Wait()
	fmt.Println("selesai")
}
