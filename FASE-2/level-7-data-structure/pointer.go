// Pointer digunakan untuk mengakses alamat data

// Value: nilai datanya.
// Pointer: nilai yang menyimpan alamat suatu variabel sehingga bisa mengakses data di alamat tersebut

package main

import (
	"fmt"
)

type Product struct {
	Name string
	Price int
	Stock int
}

// Value receiver
func (p Product) ReduceStockValue() {
	p.Stock -= 10
	fmt.Println("Stock in method ReduceStockValue:", p.Stock)
}

// Pointer receiver
func (p *Product) ReduceStockPointer() {
	p.Stock -= 10
}

func main() {
	stock := 100

	pointer := &stock

	fmt.Println("Nilai Stock:", stock)
	fmt.Println("Alamat Stok:", pointer)
	fmt.Println("Nilai melalui pointer:", *pointer)

	// Implementation Pointer
	product := Product{
		Name: "Pulpen",
		Price: 3000,
		Stock: 100,
	}

	product.ReduceStockValue()
	fmt.Println("After value receiver:", product.Stock)

	product.ReduceStockPointer()
	fmt.Println("After Pointer reciver:", product.Stock)
}