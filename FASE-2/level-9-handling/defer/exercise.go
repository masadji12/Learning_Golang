package main

import "fmt"

func stockProcess() {
	defer fmt.Println("Stock process completed")

	fmt.Println("Checking stock")
	fmt.Println("Updating stock")
}

func main() {
	fmt.Println("Start")
	stockProcess()
	fmt.Println("Back to main")
}