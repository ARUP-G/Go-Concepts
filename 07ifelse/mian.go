package main

import "fmt"

func main() {
	fmt.Println("If else loop")

	num := 3
	var meggage string

	if num > 10 {
		meggage = "Greater than 10"
	} else if num < 10 {
		meggage = "Less than 10"
	} else {
		meggage = "Equal to 10"
	}
	fmt.Println(meggage)

	if num := 20; num > 10 {
		meggage = "Greater than 10"
	} else {
		meggage = "Less than 10"
	}
	fmt.Println(meggage)

	if 9%2 == 0 {
		fmt.Println("Non-prime")
	} else {
		fmt.Println("Prime")
	}
}
