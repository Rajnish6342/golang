package main

import (
	"fmt"
	"strings"
)

func main() {
    n := 5
    for i := 1; i <= n; i++ {
        leftPart := ""
        rightPart := ""
        
        for j := 1; j <= i; j++ {
            leftPart += fmt.Sprintf("%d", j)
        }
        
        for j := i; j >= 1; j-- {
            rightPart += fmt.Sprintf("%d", j)
        }
        
        spaceCount := 2*(n-i)
        spaces := strings.Repeat(" ", spaceCount)
        
        fmt.Println(leftPart + spaces + rightPart)
    }
}
