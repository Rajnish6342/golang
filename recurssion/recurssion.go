package main

import "fmt"

func recurssion(a int) int {
	if a == 0 {
		return 0
	}
	fmt.Println(a)
	return recurssion(a - 1)
}

func main() {
	fmt.Println("Enter a number greater than 0 or 1")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Printing numbers to 1")
	recurssion(a)
}
