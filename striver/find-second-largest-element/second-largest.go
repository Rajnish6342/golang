package main

import (
	"fmt"
)
func approachSorting()  {
	a := []int{1,2,4,9,12,3,11}
	var min_index int
	var temp int
	for i := 0; i < len(a) -1; i++ {
		min_index = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min_index] {
				min_index = j
			 }
		}
		temp = a[i]
		a[i] = a[min_index]
		a[min_index] = temp
	}
	 fmt.Println("Second Smallest element", a[1])
	 fmt.Println("Second Largest element", a[len(a)-2])
}

func secondSmallest() int {
	a := []int{1,2,4,9,12,3,11}
	var small int
	var secondSmall int
	for _, v := range a {
		if(small > v){
			secondSmall = small 
			small = v
		}else if ( v < secondSmall && v != small) {
			secondSmall = v
		}
	}
	return secondSmall;
	
}

func secondLargest() int {
	a := []int{1,2,4,9,12,3,11}
	var largest int
	var secondLargest int
	for _, v := range a {
		if(largest < v){
			secondLargest = largest 
			largest = v
		}else if ( v > secondLargest && v != largest) {
			secondLargest = v
		}
	}
	return secondLargest;
	
}

func approachOptimal()  {
	 secondSmall := secondSmallest();
	 secondLargest := secondLargest();
	 fmt.Println("Second Smallest element", secondSmall)
	 fmt.Println("Second Largest element", secondLargest)
}
func main() {
	approachSorting()
}