// Panic adalah mekanisme Go untuk menghentikan alur eksekusi normal suatu program atau fungsi ketika terjadi kondisi yang yang tidak dapat ditangani
package main

import "fmt"

func main() {
	fmt.Println("Start Program")

	panic("A fatal error occured") // Terjadi kesalahan fatal

	fmt.Println("Completed Program")
}