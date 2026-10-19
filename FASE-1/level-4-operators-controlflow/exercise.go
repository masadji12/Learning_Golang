package main

import "fmt"

func main() {
	stock := 0

	if stock > 0 {
		fmt.Println("Stock Available")
	} else {
		fmt.Println("Stock Unavailable")
	}

	if stock == 0 {
		fmt.Println("Sold out")
	} else if stock <= 10 {
		fmt.Println("Safe")
	}

	status := "pending"

	switch status {
	case "pending":
		fmt.Println("Request pending approval")
	case "approved":
		fmt.Println("Request approved")
	case "rejected":
		fmt.Println("Request has rejected")
	default:
		fmt.Println("Status inactive")
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Number Ke-", i)
	}

	stk := 3

	for stk > 0 {
		fmt.Println("Stock ke-", stk)
		stk--
	}

	stok := 3

	for {
		fmt.Println("Stock ke-", stok)
		stok--

		if stok == 0 {
			break
		}
	}

	items := []string {
		"Pulpen",
		"Buku",
		"Pencil",
	}

	for _, item := range items {
		fmt.Println(item)
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		} 
		if i == 8 {
			break
		}
		fmt.Printf("Number ke-%d\n", i)
	}
}