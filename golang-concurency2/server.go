package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	sync.RWMutex
	Data string
}

func (c *Config) GetConfig() string {
	c.RLock()
	fmt.Println("Sedang membaca...")
	defer c.RUnlock()
	return c.Data
}

func (c *Config) UpdateConfig(data string) {
	fmt.Println("Mulai menulis (update)...") // Tambahkan ini
	time.Sleep(2 * time.Second)
	c.Lock()
	defer c.Unlock()
	c.Data = data
	fmt.Println("Selesai menulis!")
}

func main() {
	config := Config{Data: "Initial Config"}

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(config.GetConfig())
		}
	}()

	go config.UpdateConfig("Updated Config")

	time.Sleep(5 * time.Second)
}
