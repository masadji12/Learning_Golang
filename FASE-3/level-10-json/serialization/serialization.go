package serialization

import (
	"encoding/json"
	"fmt"
	"level-10-json/models"
)

func Serialization() {
	product := models.Product{
		ID:    1,
		Name:  "Pulpen",
		Stock: 100,
		Price: 2500,
	}

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Failed to convert to JSON:", err)
		return
	}

	fmt.Println("--- Hasil Serialization ---")
	fmt.Println(string(data))
	fmt.Println()
}