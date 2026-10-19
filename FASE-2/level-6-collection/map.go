// If slice save elemen by index, but map save data by key and value

package main

import "fmt"

func main() {
	stocks := map[string]int{
		"Pulpen": 20,
		"Buku": 15,
		"Pensil": 30,
	}

	fmt.Println("Stok pulpen:", stocks["Pulpen"])
	fmt.Println("Stok buku:", stocks["Buku"])

	stocks["Penggaris"] = 12 // Create new data
	stocks["Pulpen"] = 25 // Update data

	delete(stocks, "Buku") // Delete data

	// Pengecekan
	stock, exists := stocks["Penggaris"]

	if exists {
		fmt.Println("Stok:", stock)
	} else {
		fmt.Println("Product not found")
	}
}