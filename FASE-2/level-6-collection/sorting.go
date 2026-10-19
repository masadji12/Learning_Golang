// Searching adalah proses mengurutkan data berdasarkan aturan tertentu,
package main

import (
	"fmt"
	"sort"
)

func main() {
	// SORTING SLICE NUMBER
	stocks := []int{30, 10, 50, 20 ,40}

	sort.Ints(stocks) // Di urutkan langsung, bukan langsung menghasilkan slice baru.

	fmt.Println(stocks)

	// SORTING SLICE STRING
	products := []string{
		"Pensil",
		"Buku",
		"Pulpen",
		"Penghapus",
	}

	sort.Strings(products)

	fmt.Println(products)

	// SORT FROM LARGEST TO SMALLEST
	sort.Sort(sort.Reverse(sort.IntSlice(stocks)))

	fmt.Println(stocks)
}