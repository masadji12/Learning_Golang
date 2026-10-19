package main

import "fmt"

func main() {
	const storeName = "Anak Nunggal Store"

	var productName string
	var stock int
	var price float64
	var isActive bool = true

	fmt.Println("=== DATA BARANG ATK ===")
	fmt.Println("Toko: ", storeName)

	fmt.Println("Nama produk: ")
	fmt.Scanln(&productName) // Simbol use for mendapatkan alamat variabl

	fmt.Println("Stock: ")
	fmt.Scanln(&stock)

	fmt.Println("Harga Produk: ")
	fmt.Scanln(&price)

	totalValue := float64(stock) * price

	fmt.Println("\n=== HASIL INPUT ===")
	fmt.Printf("Nama produk : %s\n", productName)
	fmt.Printf("Stok        : %d\n", stock)
	fmt.Printf("Harga       : Rp %.2f\n", price)
	fmt.Printf("Aktif       : %t\n", isActive)
	fmt.Printf("Nilai stok  : Rp %.2f\n", totalValue)
}
