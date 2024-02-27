package test

import (
	"fmt"
	"testing"
)

// function yang akan dijalankan sebagai goroutine
func GiveMeResponse(chanString chan string, name string) {
	chanString <- name
}

func TestChannelAsParams(t *testing.T) {
	// create channel
	chanString := make(chan string)

	go GiveMeResponse(chanString, "reo")

	// consume
	data := <-chanString
	fmt.Println("value :", data)

	// test reo
	
}
