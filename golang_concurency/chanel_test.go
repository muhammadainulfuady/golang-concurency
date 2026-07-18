package golang_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello from goroutine"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("Menerima data dari channel:", data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from goroutine"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("Menerima data dari channel:", data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from goroutine"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Menerima data dari channel:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Data 1"

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChanel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- fmt.Sprintf("Data ke-%d", i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data dari channel:", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	counter := 0

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}
