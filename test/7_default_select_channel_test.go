package test

import (
	"fmt"
	"log"
	"testing"
)

func TestDefaultSelect(t *testing.T) {
	// create channel
	chanData1 := make(chan string, 1)
	chanData2 := make(chan string, 1)
	defer func() {
		close(chanData1)
		close(chanData2)
	}()

	go GiveMeResponse(chanData1, "reo")
	go GiveMeResponse(chanData2, "hello")

	var response []string
	counter := 0
	for {
		select {
		case data := <-chanData1:
			response = append(response, data)
			counter++
		case data := <-chanData2:
			response = append(response, data)
			counter++
		default:
			log.Println("waiting data")
		}

		if counter == 2 {
			break
		}
	}

	fmt.Println(response)
}
