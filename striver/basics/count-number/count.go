package main

import (
	"fmt"
	"math"
)

func main()  {
	var num  float64= 12345;
	var count int =0;
	for num > 0 {
		count ++
		num = math.Floor(num / 10)
	}
	fmt.Printf("No Of Digits %d ",count)
	
}