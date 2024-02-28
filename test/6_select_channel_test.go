package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	// create channel
	chanUser1 := make(chan string, 1)
	chanUser2 := make(chan string, 1)
	defer func() {
		close(chanUser1)
		close(chanUser2)
	}()

	// send data to channel
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		chanUser1 <- "reo"
	}()
	go func() {
		defer wg.Done()
		chanUser2 <- "sahobby"
	}()

	wg.Wait()
	counter := 1
	response := []string{}
	for {
		select {
		case user := <-chanUser1:
			response = append(response, user)
		case user := <-chanUser2:
			response = append(response, user)
		}

		if counter == 2 {
			break
		}

		counter++
	}

	// println response
	fmt.Println(response)
}
