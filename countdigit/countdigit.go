package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please enter a number: ")
	var n1 int
	fmt.Scanln(&n1)
	if n1 > 0 {
		out := noOfDigits(n1)
		fmt.Print(out)
	} else {
		fmt.Print("Please enter a number greater than zero ")
	}

}

func noOfDigits(a int) int {
	x := 0
	for a > 0 {
		a = a / 10
		x++
		noOfDigits(a)
	}
	return x
}
