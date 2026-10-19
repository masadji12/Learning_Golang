// Go memiliki fitur yang sangat berguna: satu function dapat mengembalikan lebih dari satu nilai

package main

import (
	"errors"
	"fmt"
)

func calculate(a, b int) (int, int) {
	sum := a + b
	difference := a - b

	return sum, difference
}

// Multiple return dalam validasi bisnis
// Pada backend, function sering perlu mengembalikan hasil sekaligus error.
func withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return balance, errors.New("Jumlah penarikan harus positif")
	}

	if amount > balance {
		return balance, errors.New("Saldo tidak mencukupi")
	}

	return balance - amount, nil
}

func main() {
	a := 10
	b := 3

	sum, difference := calculate(a, b)

	fmt.Printf("Jumlah %d + %d: %d\n", a, b, sum)
	fmt.Println("Selisih:", difference)

	// Multiple return dalam validasi bisnis
	balance, err := withdraw(100000, 30000)

	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	fmt.Println("Remaining balance:", balance)
}
