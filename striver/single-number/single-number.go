package main

func singleNumber(nums []int) int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		count := hashMap[v]
		hashMap[v] = count + 1
	}

	for k := range hashMap {
		if hashMap[k] == 1 {
			return k
		}
	}
	return 0
}

func main() {

}
