// Funtion not always need parameter or value resultm, function yang tidak menerima data dari luar ketika dipanggil

package main

import "fmt"

func main() {
	showWelcome()
}

func showWelcome() {
	fmt.Println("Welcome to Master ATK")
}