package main

import "fmt"

func main () {
	// Array a have ukuran tetap. Slice memberikan cara yang lebih fleksibel for working with kumpulan data
	stocks := []int{10, 20, 30}

	fmt.Println("Data stok:", stocks)
	fmt.Println("Jumlah Elemen:", len(stocks))

	// Adding element with append()
	stocks = append(stocks, 40)
	stocks = append(stocks, 50)

	fmt.Println(stocks)

	// len() dan cap()
	sizes := make([]int, 3, 5)

	fmt.Println("Long:", len(sizes)) // Jumlah elemen yang sedang ada
	fmt.Println("Capacity:", cap(sizes)) // Kapasitas slice, yaitu jumlah elemen yang dapat ditampung

	// Get sebagian elemen dengan slicing
	products := []string{
		"Pulpen",
		"Buku",
		"Pensil",
		"Penghapus",
	}

	fmt.Println(products[1:3])

	// Delete elemen dari slice
	index := 1

	products = append(products[:index], products[index+1:]...)

	fmt.Println(products)
}