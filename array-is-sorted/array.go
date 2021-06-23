package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the size of array")
	fmt.Scan(&n)
	setArray(n)
}

func setArray(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the %dth element", i)
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
	checkSorted(a, n)
}

func checkSorted(array []int, size int) {
	var check bool
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {

			if array[j] > array[i] {
				check = true
			} else {
				check = false
			}

		}
	}
	fmt.Print(check)
}
