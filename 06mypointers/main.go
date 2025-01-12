package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	println("Value is ", ptr)
	println("Value is ", *ptr)

	*ptr += 2
	println("New value is ", myNumber)

}
