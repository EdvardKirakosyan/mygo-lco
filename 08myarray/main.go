package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [7]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[6] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))
}
