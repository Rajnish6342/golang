package main

import (
	"fmt"
)

func countMaxCons(nums []int) int {
	var maxC = 0
	var currentMax = 0
	var currentVal = 0
	if len(nums) < 2 {
		return 0
	}
	for _, v := range nums {
		if v != currentVal {
			maxC = max(maxC, currentMax)
			currentVal = v
			currentMax = 1
		} else if v != 0 {
			currentMax++
		}
	}
	maxC = max(maxC, currentMax)
	return maxC
}

func main() {
	res := countMaxCons([]int{0, 0})
	fmt.Println("Max Consecetices", res)
}
