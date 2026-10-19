package main

import (
	"fmt"
	"sort"
)

func main() {
	// LINEAR SEARCH
	products := []string {
		"Pulpen",
		"Buku",
		"Pensil",
		"Penghapus",
	}

	target := "Pensil"
	foundIndex := -1

	for i, product := range products {
		if product == target {
			foundIndex = i 
			break
		}
	}

	if foundIndex != -1 {
		fmt.Println("Product found at index:", foundIndex)
	} else {
		fmt.Println("Product Not Found", foundIndex)
	}

	// BINARY SEARCH
	stocks := []int{10, 20, 30, 40, 50}

	targetBinary := 30

	index := sort.SearchInts(stocks, targetBinary)

	if index < len(stocks) && stocks[index] == targetBinary {
		fmt.Println("Data found at index:", index)
	} else {
		fmt.Println("Data not found")
	}

	// USES FOR BINARY SEARCH
	numbers := []int{10, 20, 30, 40, 50}
	targetUseFor := 40

	left := 0
	right := len(numbers) - 1
	foundIndexUseFor := -1

	for left <= right {
		mid := left + (right - left) / 2

		if numbers[mid] == targetUseFor {
			foundIndexUseFor = mid
			break
		} else if numbers[mid] < targetUseFor {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if foundIndexUseFor != -1 {
		fmt.Println("Numbers Found at index:", foundIndexUseFor)
	} else {
		fmt.Println("Number not found")
	}
}