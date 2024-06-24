package main

import (
	"fmt"
)

func union(array1, array2 []int) []int {
	unique := make(map[int]bool) // Using a map to store unique elements
	var result []int

	// Add all elements from array1 to the map
	for _, value := range array1 {
		if _, found := unique[value]; !found {
			unique[value] = true
			result = append(result, value)
		}
	}

	// Add all elements from array2 to the map
	for _, value := range array2 {
		if _, found := unique[value]; !found {
			unique[value] = true
			result = append(result, value)
		}
	}

	return result
}

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(union(array1, array2))
}
