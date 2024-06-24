package main

import "fmt"

func main() {
	a := []int{1,2,4,9,12,3,11}
	 largest := 0;
	 for _, v := range a {
			if v > largest {largest = v}
	 }
	 fmt.Println("Largest element", largest)
}