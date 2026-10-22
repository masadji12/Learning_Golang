// Defer digunakan untuk menjadwalkan pemanggilan fungsi agar dijalankan ketika fungsi yang memuatnya akan selesai

// Usually defer use for:
// - Menutup file setelah selesai dibaca
// - Close the connection or resource
// - Un lock after used
// - Menjalankan proses cleanup

package main

import "fmt"

func main() {
	fmt.Println("Start Program")

	defer fmt.Println("Message defer")

	fmt.Println("Program completed")
}