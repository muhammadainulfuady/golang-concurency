package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	sync.RWMutex // Mutex khusus untuk struct ini
	balance      int
}

// Method untuk tambah/kurang saldo (Menulis data)
func (account *BankAccount) AddBalance(amount int) {
	account.Lock()         // Pakai Lock (Write Lock) karena kita mengubah data
	defer account.Unlock() // Pastikan buka kunci setelah selesai

	account.balance += amount
	fmt.Printf("Saldo diupdate: %d\n", account.balance)
}

// Method untuk cek saldo (Membaca data)
func (account *BankAccount) GetBalance() int {
	account.RLock()         // Pakai RLock (Read Lock) karena cuma baca
	defer account.RUnlock() // Pastikan buka kunci baca

	return account.balance
}

func main() {
	account := BankAccount{}

	// Goroutine 1: Update Saldo (Write)
	for i := 0; i < 5; i++ {
		go func(amount int) {
			account.AddBalance(amount)
		}(1 * 100)
	}

	// Goroutine 2: Cek Saldo (Read)
	// Kita bisa punya puluhan/ratusan goroutine yang baca saldo secara bersamaan
	// tanpa saling menunggu!
	go func() {
		fmt.Printf("Cek saldo: %d\n", account.GetBalance())
	}()

	time.Sleep(3 * time.Second)
}
