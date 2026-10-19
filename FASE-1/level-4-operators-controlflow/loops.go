package main

import "fmt"

func main() {
	// For loop in Golang
	// for inisialisasi; kondisi; expression {}
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}

	// While loop in Golang
	// for condition {}
	stock := 3

	for stock > 0 {
		fmt.Println("In Stock:", stock)
		stock--
	}

	fmt.Println("Stock sold out")

	// Do While in Golang
	count := 1

	for {
		fmt.Println("Trying:", count)
		count++

		if count == 2 {
			continue
		} else if count == 5 {
			break
		}
	}

	// for range seperti for each
	products := []string {
		"Pulpen",
		"Buku",
		"Pensil",
	}

	for index, product := range products {
		fmt.Printf("Index Ke-%d is %s\n",index, product)
	}

	// If you dont need it index
	for _, product := range products {
		fmt.Println(product)
	}

	// Nested loop
	categories := []string{"Alat Tulis", "Kertas", "Perlengkapan IT"}

	for _, category := range categories {
		fmt.Println("Category:", category)

		for i := 1; i<= 2; i++ {
			fmt.Println("Item ke-", i)
		}
	}
}