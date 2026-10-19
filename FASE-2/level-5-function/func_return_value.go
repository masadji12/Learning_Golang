// Return value adalah nilai yang dikembalikan oleh function setelah tugasnya selesai

package main

import "fmt"

func calculateDiscount(price int, discount int) int {
	return price * discount / 100
}

func main() {
	discount := calculateDiscount(50000, 10)

	fmt.Println("Diccount:", discount)
}