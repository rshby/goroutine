package test

import (
	"fmt"
	"sync"
	"testing"
)

// function yang digunakan hanya untuk send data ke channel
func OnlyInChannel(wg *sync.WaitGroup, channel chan<- string, name string) {
	defer wg.Done()
	channel <- name
}

// function yang digunakan hanya untuk consume data dari channel
func OnlyOutChannel(wg *sync.WaitGroup, channel <-chan string) {
	defer wg.Done()
	fmt.Println(<-channel)
}

func TestChannelInOut(t *testing.T) {
	// create channel
	chanString := make(chan string, 1)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go OnlyInChannel(wg, chanString, "rere")

	wg.Add(1)
	go OnlyOutChannel(wg, chanString)

	wg.Wait()
	fmt.Println("program selesai")
}
