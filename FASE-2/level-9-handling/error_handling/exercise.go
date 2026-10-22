package main

import (
	"errors"
	"fmt"
)

func stockReduce(stock, amount int) (int, error) {
	if amount < 0 {
		return  stock, errors.New("The reduction amount must be greather than 0") // Jumlah Pengurangan harus lebih dari 0
	}
	
	if amount > stock {
		return stock, errors.New("Out of stock") // Stock tidak mencukupi
	}

	return stock - amount, nil
}

func main() {
	newStock, err := stockReduce(10, 20)

	if err != nil {
		fmt.Println("Failed:", err)
		return
	}
	
	fmt.Println("Stock after reduction:", newStock) // Stock setelah pengurangan
}