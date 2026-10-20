package main

import "fmt"

func main() {
	atks := []string{
		"Pulpen",
		"Buku",
		"Penghapus",
	}

	// Read
	fmt.Println("=== READ ===")
	for i := 0; i <= len(atks)-1; i++ {
		fmt.Println(i+1, atks[i])
	}

	// Add
	fmt.Println("\n=== ADD ===")
	atks = append(atks, "Pensil")
	fmt.Println(atks)

	// Update
	fmt.Println("\n=== UPDATE ===")
	atks[1] = "Buku Tulis"
	fmt.Println(atks)

	// Delete
	fmt.Println("\n=== DELETE ===")
	atks = append(atks[:2], atks[3:]...)
	fmt.Println(atks)
}
