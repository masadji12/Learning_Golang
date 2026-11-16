package main

import (
	"fmt"
	"sync"
)

func main() {
	// Stock awal product
	stock := 100

	// WaitGroup untuk menunggu semua goroutine
	var wg sync.WaitGroup

	// Menjalankan 5 goroutine
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()

			// Membaca stock sebanyak 10
			currentStock := stock

			// Mengurangi stock sebanyak 10
			currentStock -= 10

			// Menyimpan kembali stock
			stock = currentStock
		} ()
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	fmt.Println("Final Stock:", stock)
}