package longestconsseq

func longestConsecutive(nums []int) int {
	n := len(nums)
	l := 0
	if n == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, v := range nums {
		numSet[v] = true
	}
	for num := range numSet {
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentStreak := 1
			for {
				if _, exists := numSet[currentNum+1]; exists {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}
			if currentStreak > l {
				l = currentStreak
			}
		}
	}

	return l
}
