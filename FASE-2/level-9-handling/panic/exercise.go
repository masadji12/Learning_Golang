package main

import "fmt"

func main() {
	defer fmt.Println("Close resource")

	fmt.Println("Start Process")

	panic("Stock invalid")

	fmt.Println("Process completed")
}