package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Baris ini dieksekusi duluan secara normal
	fmt.Println("-> Fungsi utama (main) dimulai")

	// 2. Kita lempar proses ini ke background pakai 'go'
	go func() {
		fmt.Println("⏳ Goroutine mulai jalan...")

		// Simulasi proses yang makan waktu 1 detik
		time.Sleep(1 * time.Second)

		fmt.Println("✅ Goroutine beres!")
	}()

	// 3. Fungsi utama lanjut jalan tanpa harus nungguin proses di atas
	fmt.Println("-> Fungsi utama lanjut ngerjain yang lain")

	// 4. KUNCI PENTING: Kita tahan fungsi utama selama 2 detik
	// Kalau baris ini dihapus, tulisan "Goroutine beres!" ga akan pernah muncul
	// karena fungsi main keburu mati.
	time.Sleep(2 * time.Second)

	fmt.Println("-> Program selesai & ditutup")
}
