package main

import "fmt"

type Product struct {
	Name	string
	Price 	int
	Stock	int
}

func main() {
	var products [3]Product

    products[0] = Product{
        Name:  "Pulpen",
        Price: 3000,
        Stock: 100,
    }

    products[1] = Product{
        Name:  "Buku",
        Price: 5000,
        Stock: 50,
    }

    products[2] = Product{
        Name:  "Penghapus",
        Price: 2000,
        Stock: 75,
    }

	for i, product := range products {
		fmt.Println(i + 1, product.Name, product.Price, product.Stock)
	}
}