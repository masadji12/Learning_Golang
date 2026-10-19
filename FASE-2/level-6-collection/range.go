// Range adalah cara praktis untuk menelusuri setiap elemen dan array, slice, atau map.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// RANGE IN SLICE
	products := []string{"Pulpen", "Buku", "Pensil"}

	// if you don't need print index use _
	for index, product := range products {
		fmt.Println(index, product)
	}

	// RANGE IN MAP
	stocks := map[string]int{
		"Pulpen": 20,
		"Buku": 15,
		"Pensil": 30,
	}

	for product, stock := range stocks {
		fmt.Println(product, stock)
	}

	// Range tidak otomatis mengubah nilai koleksi
	// Jika ingin mengubah elemen use index

	// SORTING
	prices := []int{5000, 2000, 4000, 1000, 3000}

	sort.Ints(prices) // Mengurutkan Angka dari kecil ke besar
	fmt.Println(prices)

	sort.Strings(products) // Mengurutkan String by leksikografis (nilai byte)
	fmt.Println(products)

	sort.Sort(sort.Reverse(sort.IntSlice(prices))) // Mengurutkan dari besar ke kecil
	fmt.Println(prices)
}