package golang_concurrency

import (
	"fmt"
	"testing"
)

// pusat data kependudukan
func kirimData(nama string, channel chan string) {
	channel <- "Penduduk: " + nama
}

func TestKirimData(t *testing.T) {
	channel := make(chan string)

	go kirimData("Budi", channel)
	go kirimData("Joko", channel)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
