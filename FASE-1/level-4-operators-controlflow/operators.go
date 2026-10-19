package main

import "fmt"

func main() {
	// ARITHMETIC OPERATORS
	// +, -, *, /, %
	a := 10
	b := 3

	fmt.Println("Penjumlahan:", a + b)
	fmt.Println("Pengurangan:", a - b)
	fmt.Println("Pembagian:", a / b)
	fmt.Println("Perkalian:", a * b)
	fmt.Println("Sisa Bagi:", a % b)

	stock := 100
	sold := 15

	remainigStock := stock - sold

	fmt.Println("Initial Stock:", stock)
	fmt.Println("Soldout:", sold)
	fmt.Println("Remaining Stock:", remainigStock)

	// COMPARISON OPERATORS
	// ==, !=, >, <, >=, <=
	minimumStock := 50
	fmt.Println(stock >= minimumStock)

	// LOGICAL OPERATORS
	// &&, ||, !
	isActive := true

	canSell := stock >= 10 && isActive

	fmt.Println("The Product can be sold?", canSell)

	// ASSIGNMENT OPERATORS
	// =, +=, -=, *=, /=, ++, --
	a = 20
	b = 30

	a += b

	fmt.Printf("The result a + b = %d", a)
}