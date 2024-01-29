package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Maulana fb"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Maulana fb"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Maulanafb"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {
		channel <- "Maulana"
		channel <- "Fatih"
		channel <- "Bilqisthi"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	defer close(channel1)
	defer close(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 1 : ", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}
}
