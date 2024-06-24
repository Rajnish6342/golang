package main

import "fmt"

func moveZeroes(array [] int, size int) [] int  {
	temp := []int {}
	for _, v := range array {
		if(v != 0){
			temp = append(temp, v)
		}
	}
	for  a := 0; a < size - len(temp) ; a++ {
		temp = append(temp, 0)

	}
  return temp;	
}

func moveZeroesTwoPointer(array [] int, size int) [] int   {
	var zeroIndex = -1;
	for  a := 0; a < size  ; a++ {
		if(array[a] == 0){
			zeroIndex = a;
			break;
		}
	}
	if (zeroIndex == -1) {return array};

	for b := zeroIndex+1; b < size; b++ {
		if(array[b] != 0){
			array[zeroIndex], array[b] = array[b], array[zeroIndex]
			zeroIndex++
		}
		
	}
	return array
}
func main()  {
	var num = 7
   array := []int{2,0,0,1,0,8,5}
   fmt.Println(moveZeroes(array,num))
   fmt.Println(moveZeroesTwoPointer(array,num))
}