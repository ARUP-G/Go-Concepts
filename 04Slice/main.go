package main

import (
	"fmt"
	"sort"
)

// Slice is just a dynamic array (like array-list)
func main() {

	fmt.Println("Slice")

	//var items string

	//Slice declaration
	var shoppingList = []string{"Appel", "Mango", "Biscut"}

	//Empty slice assignment
	// var shoppingList = []string{}
	// shoppingList := []string{}

	//fmt.Scan(&items)

	shoppingList = append(shoppingList, "Banana", "Peach")

	fmt.Println("Shopping list: ", shoppingList)
	shoppingList = append(shoppingList[1:3]) //this shows from 1-2

	shoppingList = append(shoppingList[:3]) // starts from 0 as default
	fmt.Println("Portion of slice: ", shoppingList)
	fmt.Println("Portin of slice starting from 0(default): ", shoppingList)

	//memory allocation using make
	points := make([]int, 4)

	points[0] = 123
	points[1] = 823
	points[2] = 193
	points[3] = 843
	//points[4] = 903 //error as we initiated size=4

	fmt.Println("Scores: ", points)

	// but we can append in the slice as it realocates the slice in the memory
	points = append(points, 900, 750, 669, 401)
	fmt.Println("after appending: ", points)

	//sort
	sort.Ints(points)
	fmt.Println("After sorting: ", points)
	fmt.Println(sort.IntsAreSorted(points)) //boolean

	//Remove
	var index int = 2
	var names = []string{"Arup", "Rino", "Ferd", "Ritu", "Arun"}
	fmt.Println("Names were: ", names)
	names = append(names[:index], names[index+1:]...)
	fmt.Println("Without index 2: ", names)

}
