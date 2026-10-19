// Interface adalah tipe yang mendefinisikan kumpulan method yang harus dimiliki oleh suatu tipe agar memenuhi interface tersebut.

package main

import (
	"fmt"
)

type Informer interface {
	DisplayInfo()
}

type Product struct {
	Name string
	Price int
	Stock int
}

func (p Product) DisplayInfo() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.Price)
	fmt.Println("Stock:", p.Stock)
}

func main() {
	product := Product{
		Name: "Pulpen",
		Price: 3000,
		Stock: 100,
	}

	product.DisplayInfo()
}