package main

import "fmt"

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	current := 0

	for current <= right {
		if nums[current] == 0 {
			nums[current], nums[left] = nums[left], nums[current]
			left++
			current++
		} else if nums[current] == 2 {
			nums[current], nums[right] = nums[right], nums[current]
			right--
		} else {
			current++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums) // Output: [0, 0, 1, 1, 2, 2]
}
