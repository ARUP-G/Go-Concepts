package main

import "fmt"

func main() {

	fmt.Println("Pointer concepts")

	// '*' represents 2 things .
	// i)  *int -> represents type (pointer type) with base int.
	// ii) *p -> represents a variable of pointer type , here '*' acts as an operator
	//     and returns what 'p' is pointing to.

	i := 2
	p := &i
	fmt.Println("P ->", p) // Gives the address
	fmt.Printf("Type of p %T\n", p)

	fmt.Println("*p -> ", *p) // Gives the value

	// if we change *p
	*p = 8
	fmt.Println("i ->", i) // Chnages the value of 'i' through memory access

	//& -> reference (address of)
	number := 22
	var ptr = &number

	fmt.Println("number->", number)
	fmt.Println("*&number->", *&number)
	fmt.Println("&number->", &number)

	fmt.Println("Memory address:", ptr)
	fmt.Println("value inside pointer:", *ptr) // same as print(number)

}

// Benifits of pointer
// we can create a variable once and access the variable from any where.
// we can update and change the var form anyplace using the address

// new() -> allocate memory but not initialize
//			it will have a memory address
// 			zeroed storage is created ie. programmer can/t put ant data in it

// make() -> allocate memory but initialize
//			it will have a memory address
// 			non-zeroed storage is created ie. programmer can put data init

// GC -> Garbage collection happens automatically
