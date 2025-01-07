package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
	fmt.Printf("Type of rating is %T, ", input)
}
