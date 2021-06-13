package main

// Importing fmt
import (
	"fmt"
)

func main() {
	fmt.Print("Please enter first number: ")
	var n1 int
	fmt.Scanln(&n1) // take input from user
	fmt.Print("Please enter Second number: ")
	var n2 int
	fmt.Scanln(&n2) // take input from user
	fmt.Println("Sum of two numbers: ", n1+n2)
	fmt.Println("Subsctraction of two numbers: ", n1-n2)
	fmt.Println("Multiplication of two numbers: ", n1*n2)
	fmt.Println("Divison of two numbers: ", n1/n2)
	fmt.Println("Modulous of two numbers: ", n1%n2)
}
