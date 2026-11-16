// Channel adalah mekanisme komunikasi antar-goroutine untuk mengirim dan menerima data

// Misalnya we have two goroutines:
// - The first goroutine calculates the total number of transaction (Goroutine keduape)
// - The second goroutine needs to perform than calculation (Goroutine kedua perlu menghasilkan perhitungan tersebut)

// The Channel have data type. Thats means channel bertipe int send dan receive	data type int.
package main

import (
	"fmt"
)

func main() {
	// FUNDAMENTAL CHANNEL
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	message := <-ch

	fmt.Println(message)

	// UNBUFFERED CHANNEL -> Does not have temporaty storage
	// (Tidak memiliki tempat penyimpanan sementara)
	chub := make(chan int)

	go func() {
		chub <- 100
		fmt.Println("The data was successfully")
	}()

	result := <-chub
	fmt.Println("Data diterima:", result)

	// BUFFERED CHANNEL -> Has the capacity to storage temporary data 
	// (Memiliki kapasitas untuk menampung data sementara)
	chb := make(chan int, 3)

	chb <- 100
	chb <- 200
	chb <- 300

	fmt.Println(<-chb)
	fmt.Println(<-chb)
	fmt.Println(<-chb)
}