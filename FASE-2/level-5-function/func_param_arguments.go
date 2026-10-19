// Parameter adalah variabel yang didefinisikan pada function. Argument adalah nilai yang diberikan ketika function dipanggil.

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

// Saveral parameter of the same type
func add(a, b int) int {
	result := a + b
	return result
}

func main() {
	greet("Adji Gans")
	greet("Masadji Gans")

	// Parameter of the same type
	add := add(1, 2)
	fmt.Println(add)
}