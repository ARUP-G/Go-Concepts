package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array")

	var fruitList [3]string

	fruitList[0] = "Mango"
	fruitList[2] = "Bananna"

	fmt.Scan(&fruitList[1])

	fmt.Println("Fruit list: ", fruitList)
	fmt.Printf("Fruit list:%v \n", fruitList)

	fmt.Println("Length of array", len(fruitList)) // always will be the same as declared
}
