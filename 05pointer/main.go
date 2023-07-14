package main

import "fmt"

func main() {

	fmt.Println("Pointer concepts")

	// var poin *int
	// fmt.Println("Value of Pointer :", poin)

	//& -> reference
	number := 22
	var ptr = &number

	fmt.Println("number->", number)
	fmt.Println("*&number->", *&number)
	fmt.Println("&number->", &number)

	fmt.Println("Memory address:", ptr)
	fmt.Println("value inside pointer:", *ptr) // same as print(number)

}

// new() -> allocate memory but not initialize
//			it will have a memory address
// 			zeroed storage is created ie. programmer can/t put ant data in it

// make() -> allocate memory but initialize
//			it will have a memory address
// 			non-zeroed storage is created ie. programmer can put data init

// GC -> Garbage collection happens automatically
