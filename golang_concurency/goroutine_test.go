package golang_concurrency

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Mengetes goroutine dan menjalankan fungsi RunHelloWorld")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Creating a goroutine")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Menampilkan angka:", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
