// Concurrency adalah kemampuan program untuk menangani beberapa pekerjaan dalam satu periode waktu.

// Parallelisme adalah kondisi ketika beberapa pekerjaan benar-benar dijalankan secara bersamaan, usually menggunakan beberapa CPU core.

// Goroutine adalah unit eksekusi ringan (lightweight execution unit) yang dikelola oleh goroutine. dengan goroutine, we can running function tanpa harus menunggu fungsi selesai terlebih dahulu.

package main

import (
	"fmt"
	"time"
)

func downloadImage() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Download image:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func sendEmail() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Send Email:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func generateReport() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Generate Report:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go downloadImage()
	go sendEmail()
	go generateReport()

	// time.Sleep() digunakan karena kita perlu memberikan waktu kepada goroutine agar dapat menyelesaikan pekerjaannya sebelum program utama berakhir(main function berakhir)
	time.Sleep(5 * time.Second)

	fmt.Println("All tasks finished")
}