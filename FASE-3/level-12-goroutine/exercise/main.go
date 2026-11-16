// Combining goroutine, WaitGroup, dan Channel
// Misalnya backend toko menjalankan tiga pekerjaan untuk menghitung jumlah stok dari tiga gudang.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Memproses data product
func processProduct(wg *sync.WaitGroup, ch chan string) {
	// Done() mengurangi counter WaitGroup sebanyak 1
	// defer memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Product processing started")

	// Sleep() menghentikan sementara goroutine selama 1 detik
	time.Sleep(1 * time.Second)

	// Mengirim pesan ke channel
	ch <- "Product processing completed"
}

// Memproses data order
func processOrder(wg *sync.WaitGroup, ch chan string) {
	// Done() mengurangi coutner WaitGroup sebanyak 1
	// defer memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Product processing started")

	// Sleep() menghentikan sementara goroutine selama 1 detik
	time.Sleep(1 * time.Second)

	// Mengirim pesan ke channel
	ch <- "Order processing completed"
}

func processStock(wg *sync.WaitGroup, ch chan string) {
	// Done() mengurangi counter WaitGroup sebanyak 1 kali
	// defer memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Stock processing started")

	// Sleep() menghentikan sementara goroutine selama 1 detik
	time.Sleep(1 * time.Second)

	// Mengirimi pesan ke channel
	ch <- "Stock processing completed"
}

func main() {
	// Membuat WaitGroup untuk mengatur tiga goroutine
	var wg sync.WaitGroup

	// Membuat buffered channel dengan capasity 3
	// channeil ini dapat menampung maksimal 3 pesan sebelum dibaca
	ch := make(chan string, 3)

	// Add(3) mendaftarkan tiga pekerjaan ke WaitGroup
	wg.Add(3)

	// Menjalankan ketiga fungsi secara bersamaan menggunakan goroutine
	go processProduct(&wg, ch)
	go processOrder(&wg, ch)
	go processStock(&wg, ch)

	// Wait() menunggu sampai semua goroutine selesai
	wg.Wait()

	// close() menandakan bahwa tidak ada lagi data yang akan dikirim ke channel
	close(ch)

	// range membaca setiap pesan dari channel sampai channel ditutup
	for message := range ch {
		fmt.Println(message)
	}

	fmt.Println("All backend tasks completed")
}