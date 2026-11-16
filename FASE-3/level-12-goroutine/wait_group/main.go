// sync.WaitGroup adalah fasilitas dari package sync untuk menunggu sekumpulan goroutine sampai semuanya selesai

package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadImage(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Download image...")
	time.Sleep(2 * time.Second)
	fmt.Println("Image downloaded!")
}

func sendEmail(wg *sync.WaitGroup) {
	defer wg.Done() // Will be callen when the function finishes

	fmt.Println("Sending email...")
	time.Sleep(1 * time.Second)
	fmt.Println("Email sent!")
}

func generateReport(wg *sync.WaitGroup) {
	defer wg.Done() // Will be called when the function finishes

	fmt.Println("Generating report...")
	time.Sleep(3 * time.Second)
	fmt.Println("Report generated!")
}

func main() {
	// Creating a WaitGroup
	var wg sync.WaitGroup

	// Registration work 3
	wg.Add(3)

	// Run a goruotine
	go downloadImage(&wg) // Sending a pointer WaitGroup
	go sendEmail(&wg)
	go generateReport(&wg)

	// Waiting for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks finished")
}