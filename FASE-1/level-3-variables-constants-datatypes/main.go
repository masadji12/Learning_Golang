package main

import "fmt"

func main() {
	// the only var use for declaration variable data
	// wajib digunakan jika mendeklarasikan variable di luar fungsi global
	var productName string = "Pulpen"
	var stock int = 100

	fmt.Println(productName)
	fmt.Println(stock)

	// ## SHORTHAND DECLARATION
	name := "Buku Tulis"
	price := 4500.0

	fmt.Println(name)
	fmt.Println(price)

	// ## KONSTANTA (const) nilai const tidak bisa diubah (menyimpan nilai tetap)
	const minimumStock = 10
	fmt.Println("Minimum Stock Level", minimumStock)

	// ## TIPE DATA
	// int (bilangan bulat exm: 1. deft: 0)
	// float64 (bilangan desimal exm: 1.5, deft: 0)
	// string (text, deft:String kosong)
	// bool (boolean true or false, deft: false)

	// ## ZERO VALUES (Nilai awal otomatis yang diberikan GO)
	var isActvie bool

	fmt.Println(isActvie)

	// ## Type Conversion (Proses mengubah nilai dari satu ke tipe lain)
	priceInt := int(price)

	fmt.Println(priceInt)

	// ## Menampilkan data dengan fmt.Printf
	fmt.Printf("Produk: %s\n", name)

	// %s = String
	// %d = Bilangan bulat
	// %f = Bilangan desimal
	// %.2f = Desimal dengan 2 angka di belakang koma
	// %t = Boolean
	// %T = Menampilkan tipe data
	// %v = Format umum sesuai nilai
	// \n = pindah ke baris berikutnya
}