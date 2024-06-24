package main

import (
	"fmt"
)

func main() {
	num := 12345
	var reversedNum int

	for num > 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}

	fmt.Printf("Reversed Number: %d\n", reversedNum)
}
