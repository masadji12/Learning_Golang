// Struct adalah tipe data yang mengelompokkan beberapa field dengan tipe data berbeda ke dalam satu struktur

package main

import "fmt"

// type = inisialisasi tipe data baru yaitu Product
// struct = tipe tersebut berisi kumpulan field
type Product struct {
	Name	string
	Price	int
	Stock	int
}

func main() {
	product := Product{
		Name: "Pulpen",
		Price: 3000,
		Stock: 100,
	}

	fmt.Println(product)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.Stock)
}