// recover() adalah fungsi bawaan go yang digunakan untuk menangani panic agar tidak terus menyebar dan menghentikan program. However, ada aturan penting: recover harus dipanggil dari fungsi yang dijalankan melalui defer agar dapat menangkap panic yang sedang berlangsung

package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic successfully handled:", r)
		}
	} ()

	fmt.Println("Start Process")

	panic("Stock invalid")

	fmt.Println("Process completed")
}