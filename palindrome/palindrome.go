package main

import "fmt"

func palindrome(num int) bool {
	input_num := num
	var remainder int
	res := 0
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	return input_num == res
}

func main() {

	fmt.Print("Enter a number to check whether it is palindrome or not \n")
	var num int
	fmt.Scanln(&num)

	var result bool = palindrome(num)
	fmt.Print(result)

}
