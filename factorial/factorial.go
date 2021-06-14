package main

import "fmt"

func main() {
	var res int16 = 1
	fmt.Print("Enter a number to find factorial \n")
	var num int
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		res = res * i
	}
	fmt.Print(res)

}
