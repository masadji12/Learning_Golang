// Map adalah struktur data yang menyimpan pasangan key dan value, untuk mengakses data menggunakan key.

package main

import "fmt"

func main() {
	// LITERAL MAP
	fmt.Println("=== LITERAL MAP ===")
	stocks := map[string]int{
		"Pulpen": 100,
		"Buku": 50,
		"Penghapus": 75,
	}

	fmt.Println(stocks)

	// MAKE MAP KOSONG EMPTY
	stock := make(map[string]int)

	// CREATE
	fmt.Println("\n=== CREATE ===")
	stock["Pulpen"] = 100
	stock["Buku"] = 50
	
	fmt.Println(stock)

	// READ
	fmt.Println("\n=== READ ===")
	fmt.Println(stock["Pulpen"])

	for name, total := range stock {
		fmt.Printf("%s %d\n",name, total)
	}

	// HOW TO IF THE KEY NOT FOUND
	fmt.Println("\n=== HOT TO IF THE KEY NOT FOUND ===")
	fmt.Println(stock["Spidol"]) // Return Zero Value (0)

	total, ok := stock["Spidol"]

	if ok {
		fmt.Println("Stock:", total)
	} else {
		fmt.Println("Product Not Found")
	}

	// UPDATE
	fmt.Println("\n=== UPDATE ===")
	stock["Pulpen"] = 150

	fmt.Println(stock["Pulpen"])

	// DELETE 
	delete(stock, "Buku")
}