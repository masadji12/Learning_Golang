package main

import "fmt"

type Category struct {
	Name	string
}

type Product struct {
	Name	string
	Stock	int
	Category
}

type Supplier struct {
	Name	string
	Telephone	string
}

func (p Product) DisplayInfo() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Stock:", p.Stock)
	fmt.Println("Category:", p.Category.Name)
}

func (s Supplier) DisplayInfo() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Telephone:", s.Telephone)
}

type Informer interface{
	DisplayInfo()
}

func printInfo(item Informer) {
	item.DisplayInfo()
}

func main() {
	product := Product{
		Name: "Pensil",
		Stock: 100,
		Category: Category{
			Name: "Alat Tulis",
		},
	}

	supplier := Supplier{
		Name: "Yanto",
		Telephone: "089089789854",
	}

	printInfo(product)
	fmt.Println("----")
	printInfo(supplier)
}