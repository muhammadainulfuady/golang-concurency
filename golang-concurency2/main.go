package main

import (
	"fmt"
	"time"
)

func main() {
	serverA := make(chan string)
	serverB := make(chan string)
	fmt.Println("Menunggu data.....")

	go func() {
		time.Sleep(5 * time.Second)
		serverA <- "Data dari server A"
	}()

	go func() {
		time.Sleep(6 * time.Second)
		serverB <- "Data dari server B"
	}()

	select {
	case data := <-serverA:
		fmt.Println("Data diterima : ", data)
	case data := <-serverB:
		fmt.Println("Data diterima : ", data)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! Pencarian dibatalkan.")
	default:
		fmt.Println("Tidak ada data yang diterima.")
	}
	time.Sleep(10 * time.Second)
}
