package deserialization

import (
	"encoding/json"
	"fmt"
	"level-10-json/models"
)

func Deserialization() {
	jsonData := []byte(`{
		"id": 1,
		"name": "Pulpen",
		"stock": 100
	}`)

	var product models.Product

	err := json.Unmarshal(jsonData, &product)
	if err != nil {
		fmt.Println("Failed to read JSON:", err)
		return
	}

	fmt.Println("--- Hasil Deserialization ---")
	fmt.Println("ID:", product.ID)
	fmt.Println("Nama:", product.Name)
	fmt.Println("Stok:", product.Stock)
	fmt.Println("Harga:", product.Price)
}