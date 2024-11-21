package main

func maxSubArray(nums []int) int {
	maxi := -9 * 10000
	sum := 0
	for _, v := range nums {
		sum += v
		maxi = max(maxi, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxi
}
