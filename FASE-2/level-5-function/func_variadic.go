// Function yang dapat menerima nol, satu, atau banyak argument untuk satu parameter variadic

package main

import "fmt"

func sumPrices(prices ...int) int {
	total := 0

	for _, price := range prices {
		total += price
	}

	return total
}

func main() {
	fmt.Println(sumPrices(1000, 2000))
	fmt.Println(sumPrices(1000, 2000, 3000))
	fmt.Println(sumPrices(sumPrices()))
}