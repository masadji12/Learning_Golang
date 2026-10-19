// Composition berarti membuat suatu tipe dengan menggabungkan atau menggunakan beberapa tipe lain.

package main

import (
	"fmt"
)

type Category struct {
	Name string
}

type Product struct {
	Name	string
	Price 	int
	Stock 	int
	Category Category
}

func main() {
	product := Product{
		Name: "Pulpen",
		Price: 3000,
		Stock: 100,
		Category: Category{
			Name: "Alat Tulis",
		},
	}

	fmt.Println()
}