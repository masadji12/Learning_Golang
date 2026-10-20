package main

import "fmt"

func main() {
	// LITERAL MAP
	stock := map[string]int{
		"Pulpen": 100,
		"Buku": 50,
		"Penghapus": 75,
	}

	// CREATE
	stock["Pensil"] = 60
	
	// READ
	fmt.Println(stock["Pulpen"])
	fmt.Println("")

	// UPDATE
	delete(stock, "Buku")

	// READ ALL
	for key, name := range stock {
		fmt.Printf("%s %d\n", key, name)
	}

	// KEY NIL
	total, ok := stock["Pulpen"]

	if ok {
		fmt.Println("\nTotal Stock:", total)
	} else {
		fmt.Println("\nKey not found")
	}
}