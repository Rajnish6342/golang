package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = 1 + i*3
	}
	fmt.Print(arr)
}
