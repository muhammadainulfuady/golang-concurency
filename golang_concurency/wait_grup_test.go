package golang_concurrency

import (
	"fmt"
	"sync"
	"testing"
)

type Stats struct {
	sync.RWMutex
	Count int
}

func (stats *Stats) Increment() {
	stats.Lock()
	defer stats.Unlock()
	stats.Count++
}

func (stats *Stats) GetValue() int {
	stats.RLock()
	defer stats.RUnlock()
	return stats.Count
}

func TestLogStats(t *testing.T) {
	var wg sync.WaitGroup
	stats := Stats{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stats.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count:", stats.GetValue())
}

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(OnlyOnce)
			// OnlyOnce()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
} 
