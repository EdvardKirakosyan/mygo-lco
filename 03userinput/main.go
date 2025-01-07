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

}
