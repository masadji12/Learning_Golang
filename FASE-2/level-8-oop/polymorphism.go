package main

import "fmt"

type Informer interface {
	DisplayInfo()
}

type Product struct {
	Name	string
	Stock	int
}

func (p Product) DisplayInfo() {
	fmt.Println("Product:", p.Name)
	fmt.Println("Stok:", p.Stock)
}

type Supplier struct {
	Name 	string
	Phone 	string
}

func (s Supplier) DisplayInfo() {
	fmt.Println("Supplier", s.Name)
	fmt.Println("Telphone:", s.Phone)
}

// Function menerima interface, bukan struct tertentu
func printInfo(item Informer) {
	item.DisplayInfo()
}

func main() {
	product := Product{
		Name: "Pulpen",
		Stock: 100,
	}

	supplier := Supplier{
		Name: "PT ATK Jaya",
		Phone: "085965895803",
	}

	printInfo(product)
	fmt.Println("---")
	printInfo(supplier)
}