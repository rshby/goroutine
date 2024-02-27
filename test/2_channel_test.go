package test

import (
	"fmt"
	"log"
	"testing"
)

/**
- channel adalah tempat komunikasi secara synchronus yang bisa dilakukan oleh goroutine
- di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
- saat melakukan pengiriman data ke channel, goroutine akan terblock, sampai ada yang menerima data tersebut
*/

/*
untuk mengirim data ke channel = channel <- data
untuk consume data dari channel = variabel <- channel
*/

func TestCreateChannel(t *testing.T) {
	// create channel untuk menyimpan data string
	chanString := make(chan string)
	defer close(chanString)

	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	// send data ke channel
	go func() {
		//defer wg.Done()
		chanString <- "reo sahobby"
		log.Println("selesai mengirim data ke channel")
	}()

	//wg.Wait()

	// consume and println value from channel
	fmt.Println("value :", <-chanString)
}
