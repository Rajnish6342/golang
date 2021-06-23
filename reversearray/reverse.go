package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter size of array")
	fmt.Scan(&size)

	array := make([]int, size)
	revered := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %d elemt", i)
		fmt.Scan(&array[i])
	}
	
	fmt.Println(revered)
}
