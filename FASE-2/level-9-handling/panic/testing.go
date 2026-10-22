package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	} ()

	fmt.Println("Start process")

	panic("Stock invalid")

	fmt.Println("Process completed")
}