// Visibility adalah aturan yang menentukan apakah sebuah nama identifier seperti fungsi, variable, konstanta, atau tipe data dapat diakses dari package lain.

// - Huruf kapital diawal nama -> exported, bisa diakses dari package lain
// - Huruf kecil diawal name -> unexported, hanya bisa diakses dari package yang sama

// Pascal Case = PascalCase
// Camel Case = camelCase
// Snake Case = snack_case
// Kebab Case = kebab-case

package main

var ProductName = "Pulpen"
var stock = 100

func DisplayProduct() string {
	return ProductName
}

func calculateStock() int {
	return stock
}

