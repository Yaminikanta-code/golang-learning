package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string //compulsory to provide no of elements
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	//prints [Apple Tomato  Peach] longer gap between Tomato and Peach because there is one empty string at fruitList[2]

	fmt.Println("Length of fruitList is: ", len(fruitList)) //always prints 4

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println(vegList[1])

}
