package main

import (
	"fmt"
)

func main() {
	fmt.Println("Function module")
	printMessage()
	fmt.Println("Is prime? ->", primeNumber(9))

	result := massAdder(2, 4, 2, 1, 8, 5)
	fmt.Println("result:", result)

	// calling massAdder2 (1. with another var)
	result3, message := massAdder2(2, 5, 2, 1, 5)
	fmt.Println("Result3: ", result3)
	fmt.Println("Message:", message)

	// calling massAdder2 (2. with comma ok)
	// won't get the message
	result2, _ := massAdder2(2, 5, 2, 1, 5)
	fmt.Println("Result2: ", result2)

}

func printMessage() {
	fmt.Println("Heeeee")
}

// bool is a signeture of the function / return type
func primeNumber(a int) bool {
	if a%2 == 0 {
		return false
		//fmt.Println("Non-Prime")
	} else {
		return true
		//fmt.Println("Prime")
	}
}

// '...' -> veriadic function
func massAdder(values ...int) int {
	total := 0
	for _, val := range values { //here '_' is index of loop
		total += val
	}
	return total
}

// func with 2 signetures
// syntax
// func name_of_functon (name type-of_parameter) type_of_function
func massAdder2(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Result is printed"
}
