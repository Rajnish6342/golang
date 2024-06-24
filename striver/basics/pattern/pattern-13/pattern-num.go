package main

import "fmt"

func main() {
    for j := 1; j <= 7; j++ {
        for i := 4; i >= 1; i-- {
            fmt.Print(i)
         for k := i-1; k <i ; k++ {
            fmt.Print(k)
         }
         
        }
        fmt.Println()
    }
   
}
