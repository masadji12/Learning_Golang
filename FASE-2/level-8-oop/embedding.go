package main

import "fmt"

type Category struct {
    Name string
}

func (c Category) DisplayCategory() {
    fmt.Println("Kategori:", c.Name)
}

type Product struct {
    Name string
    Category // embedding
}

func main() {
    product := Product{
        Name: "Pulpen",
        Category: Category{
            Name: "Alat Tulis",
        },
    }

    fmt.Println("Barang:", product.Name)
    fmt.Println("Kategori:", product.Category.Name)

    // Method dari embedded type dapat dipanggil melalui Product
    product.DisplayCategory()
}