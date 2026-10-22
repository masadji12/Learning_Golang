// Error handling is cara menangani kondisi ketika sebuah operasi gagal. Dalam Go, kesalahan yang diperkirakan terjadi biasanya ditangani menggunakan nilai bertipe error, bukan dengan try-catch.

package main

import (
	"errors"
	"fmt"
)

func checkStock(stock int) error {
	if stock < 0 {
		return errors.New("Stock can't be negative") // Stock tidak boleh negative
	}

	return nil
}

func main() {
	err := checkStock(-5)

	if err != nil {
		fmt.Println("An error occured:", err) // Terjadi kesalahan
		return
	}

	fmt.Println("Stock valid")
}