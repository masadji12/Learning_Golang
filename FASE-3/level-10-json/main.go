package main

import (
	"fmt"
	"level-10-json/deserialization"
	"level-10-json/serialization"
)

func main() {
	fmt.Println("=== Memulai Program JSON Go ===")
	fmt.Println()

	serialization.Serialization()
	deserialization.Deserialization()
}