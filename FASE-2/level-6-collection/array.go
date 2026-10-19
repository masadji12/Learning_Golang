// Array adalah kumpulan elemen dengan tipe data yang sama dan jumlah elemen yang tetap.

package main

import "fmt"

func main() {
	// How to make array
	// var arrayName [quantityElemen]dataType
	var stocks [3]int

	stocks[0] = 20
	stocks[1] = 15
	stocks[2] = 30

	for i := 0; i <= len(stocks) - 1; i++ {
		fmt.Println(i, stocks[i])
	}

	fmt.Println(stocks)

	// Creating array with literally value
	prices := [3]int{2500, 5000, 3000}

	fmt.Println(prices[0])

	// Menghitung jumlah elemen secara otomatis
	sizes := [...]int{25, 100, 50}

	fmt.Println(len(sizes))

	// Mengubah nilai array
	sizes[0] = 30

	fmt.Println(sizes[0])

	// Merge array with for
	for i := 0; i <= len(sizes) - 1; i++ {
		if sizes[1]  {
		}
		fmt.Println("Index:", i, "| Sizes:", sizes[i])
	}
}
