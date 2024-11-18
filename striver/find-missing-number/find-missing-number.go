package main

import "fmt"

func missingNumber(nums []int, size int) int {
	var target int = (size * (size + 1)) / 2
	var cur int = 0
	for _, v := range nums {
		cur += v
	}
	fmt.Println(target, cur)
	return target - cur
}

func main() {
	res := missingNumber([]int{1, 2, 3, 5}, 5)
	fmt.Println("Missing Number is", res)
}
