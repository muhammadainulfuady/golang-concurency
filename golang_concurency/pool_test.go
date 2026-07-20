package golang_concurrency

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	var wg sync.WaitGroup

	pool.Put("Hello, Pool!")
	pool.Put("Golang Concurrency")
	pool.Put("sync.Pool Example")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Pool test completed.")
}

type Data struct {
	Value string
}

func TestPoolWithStruct(t *testing.T) {
	var wg sync.WaitGroup
	pool := sync.Pool{
		New: func() any {
			fmt.Println("--- Membuat objek baru ---")
			return &Data{Value: ""}
		},
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data := pool.Get().(*Data)
			data.Value = fmt.Sprintf("Data dari Goroutine ke-%d", i)
			fmt.Println(data.Value)
			data.Value = ""
			pool.Put(data)
			time.Sleep(3 * time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Println("Pool with struct test completed.")

}

type User struct {
	Name string
}

func TestSyncPoolLoop(t *testing.T) {

	userPool := sync.Pool{
		New: func() any {
			t.Log("Membuat User baru")

			return &User{}
		},
	}

	for i := 1; i <= 5; i++ {

		user := userPool.Get().(*User)

		user.Name = "User" + strconv.Itoa(i)

		t.Log("Request", i, user.Name)

		// reset data sebelum dikembalikan
		user.Name = ""

		userPool.Put(user)
	}

}
