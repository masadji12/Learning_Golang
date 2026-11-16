package main

import (
	"fmt"
	"sync"
	"time"
)

func processProduct(wg *sync.WaitGroup, ch chan string) {
	// Done mengurangi counter WaitGroup sebanyak 1
	// Defer memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Product processing started")

	time.Sleep(1 * time.Second)

	ch <- "Product processing completed"
}

func processOrder(wg *sync.WaitGroup, ch chan string) {
	// Done untuk mengurangi counter WaitGroup sebanyak 1
	// Defer untuk memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Order processing started")

	// Menghentikan sementara goroutine selama 1 detik
	time.Sleep(1 * time.Second)

	// Mengirim pensan ke channel
	ch <- "Order processing completed"
}

func processStock(wg *sync.WaitGroup, ch chan string) {
	// Done untuk mengurangi counter WaitGroup sebanyak 1
	// Defer untuk memastikan Done() dipanggil saat fungsi selesai
	defer wg.Done()

	fmt.Println("Stock processing started")

	// Menghentikan sementara goroutine selama 1 detika
	time.Sleep(1 * time.Second)

	// Mengirin pesan ke channel
	ch <- "Stock processing completed"
}

func main() {
	// Membuat WaitGroup untuk mengatur goroutine
	var wg sync.WaitGroup

	// Membuat buffered channel dengan 3 kapasitas
	// Channel ini dapat menampung maksimal 3 pesan sebelum dibaca
	ch := make(chan string, 3)

	// Mendaftarkan ketiga pekerjaan ke WaitGroup
	wg.Add(3)

	// Menjalankan ketiga fungsi secara bersamaan menggunakan goroutine
	go processProduct(&wg, ch)
	go processOrder(&wg, ch)
	go processStock(&wg, ch)

	// Menunggu sampa goroutine selesai semua
	wg.Wait()

	// Menandakan bahwa tidak ada lagi data yang akan dikirim kechannel
	close(ch)

	// range membaca setiap pesan dari channel sampai ditutup
	for message := range ch {
		fmt.Println(message)
	}

	fmt.Println("All backend tasks completed")
}