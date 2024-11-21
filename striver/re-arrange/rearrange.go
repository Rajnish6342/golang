package rearrange

func replaceElements(arr []int) []int {
	if len(arr) == 1 {
		return [-1]
	}
	n := len(arr)
	i := n - 2
	if i >= 0 {
		j := n - 3
		for arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr;
}
