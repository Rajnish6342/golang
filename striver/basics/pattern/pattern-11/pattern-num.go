package main

import "fmt"

func main()  {
	var n int = 5;
	start := 1
	for i := 1; i <= n; i++ {
		if i % 2 == 0 {
			start =0
		} else {
			start = 1
		}
		for j := 0; j < i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}