package main

import "fmt"

type Product struct {
	Name string
	Price int
	Stock int
}

func (p Product) TotalValue() int {
	return p.Price * p.Stock
}

func main() {
	product := Product{
		Name: "Pulpen",
		Price: 3000,
		Stock: 100,
	}

	fmt.Println("Name:", product.Name)
	// product in below is receiver
	fmt.Println("Nilai Stock:", product.TotalValue())
}