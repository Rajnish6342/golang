package main

import "fmt"

func main() {
	var size int
	var search int

	fmt.Println("Enter size of array")
	fmt.Scan(&size)
	var array = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter input at %d position", i)
		fmt.Scan(&array[i])
	}

	fmt.Print("Enter element to search")
	fmt.Scan(&search)
	var res = binarySearch(array, size, search)
	fmt.Print(res)
}

func binarySearch(array []int, size int, search int) []int {
	var res []int
	for i := 0; i < size; i++ {

		if array[i] == search {
			res = append(res, i)
		}
	}
	return res
}

/*
 Another approach  divide the low and high
 to get the midpoint and check if
 searching elemnt is less or greather
 if mid== search return
 if mid < search then change the low to mid + 1
 if mid > search then change the high to mid-1
*/
