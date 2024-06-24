package main

import "fmt"

func main()  {
	var n int = 5;
	for i := 1; i <= n; i++ {
		for j := 0; j < (n - i); j++ {
			fmt.Print(" ")
		}
		for k:=0; k<(2* i -1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 5; i >= 1; i-- {
		for j := 0; j < (n - i); j++ {
			fmt.Print(" ")
		}
		for k:=0; k<(2* i -1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}