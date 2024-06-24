package main

import (
	"fmt"
)

func linearSearch(array []int, size int, item int) int {
	var start int = 0
	var end int = size - 1
	if array[start] == item {
		return start
	}
	for start < end {
		cur := start + (end-start)/2
		if array[cur] == item {
			return cur
		}
		if array[cur] > item {
			end = cur - 1
		} else {
			start = cur + 1
		}
	}
	return -1
}

func main() {
	var num = 7
	var array = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(linearSearch(array, num, 4))
}
