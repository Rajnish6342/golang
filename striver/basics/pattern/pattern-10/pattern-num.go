package main

import "fmt"

func main()  {
	var n int = 5;
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
		fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 5; i >= 1; i-- {
		for j := 0; j < i; j++ {
		fmt.Print("*")
		}
		fmt.Println()
	}
}