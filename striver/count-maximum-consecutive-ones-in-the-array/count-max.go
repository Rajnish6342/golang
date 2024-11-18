package main

import (
	"fmt"
)

func countMaxCons(nums []int) int {
	var maxC = 0
	var currentMax = 0
	var currentVal = 0
	for _, v := range nums {
		if v != currentVal {
			maxC = max(maxC, currentMax)
			currentVal = v
			currentMax = 1
		} else {
			currentMax++
		}
	}
	maxC = max(maxC, currentMax)
	return maxC
}

func main() {
	res := countMaxCons([]int{1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1})
	fmt.Println("Max Consecetices", res)
}
