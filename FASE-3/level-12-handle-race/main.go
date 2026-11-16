package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initial stock
	stock := 100

	// Membuat WaitGroup untuk mengatur goroutine
	var wg sync.WaitGroup

	// Melindungin akses terhadap stock
	var mu sync.Mutex

	// Mendaftarkan 5 pekerjaan 
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()

			// Mengunci hak akses kebagian kode dibawah ini
			mu.Lock()

			currentStock := stock

			currentStock -= 10

			stock = currentStock

			mu.Unlock()
		} ()
	}

	wg.Wait()

	fmt.Println("Final Stock:", stock)
}