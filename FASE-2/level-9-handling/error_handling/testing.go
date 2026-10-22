package main

import (
	"errors"
	"fmt"
)

func checkStock(stock int) (int, error) {
	if stock < 0 {
		return stock, errors.New("Stock can't be negative")
	}

	return stock, nil
}

func main() {
	stock, err := checkStock(-1)

	if err != nil {
		fmt.Println("An error occured:", err, stock)
		return
	}

	fmt.Println("Stock valid")
}