package subarray

func subarraySum(nums []int, k int) int {
	cCount := 0
	cSum := 0
	sSum := 0
	for _, v := range nums {
		acc := v + cSum
		if acc < k {
			cCount++
			cSum += v
		} else if acc == k {
			sSum = max(cSum, sSum)
			cSum = v
			cCount = 1
		} else {
			cCount = 1
			cSum = v
		}
	}
	return sSum
}
