// Function adalah blok kode yang dibuat untuk menjalankan tugas tertentu dan dapat dipanggil kembali ketika dibutuhkan

package main

import "fmt"

func main() {
	// For Example Without Function
	stock := 10
	price := 5000

	total := stock * price

	fmt.Println("Without Function", total)

	// For Example wiht Function
	funcTotal := calculateTotal(stock, price)
	fmt.Println("Use Function: ", funcTotal)
}

// For Example with Function and Parameter
func calculateTotal(stock int, price int) int {
	return stock * price
}