// if, else if, else
// Memilih tindakan berdasarkan kondisi. Contoh: menentukan apakah stok habis, menipis, atau aman.

// switch
// Memilih tindakan berdasarkan nilai atau kondisi. Contoh: menentukan tindakan berdasarkan status barang

// for, break, continue
// Mengulan proses (for), menghentikan loop (break), or melewati iterasi (continue)

package main

import "fmt"

func main() {
	isActive := true
	stock := 10 
	requestedQty := 5

	// if
	if stock > 0 {
		fmt.Println("Stock Available")
	}

	// if-else
	if stock < 0 {
		fmt.Println("Stock Unavailable")
	} else {
		fmt.Println("Stock Available")
	}

	// if-else if-else
	if stock == 0 {
		fmt.Println("Stock sold out")
	} else if stock <= 10 {
		fmt.Println("Stock low")
	} else {
		fmt.Println("Stock Available")
	}

	// Nested if
	if isActive {
		if stock >= requestedQty {
			fmt.Println("Request can been the process")
		} else {
			fmt.Println("Insufficient Stock")
		}
	} else {
		fmt.Println("Product inactive")
	}

	// switch
	status := "approved"

	switch status {
	case "pending":
		fmt.Println("Request pending approval")
	case "approved":
		fmt.Println("Request approved")
	case "rejected":
		fmt.Println("Request denied")
	default:
		fmt.Println("Unknown status")
	}

}