package main

import "fmt"

// defer -> key word works in LIFO format
// defer puts itself at the last of the main func line after all non-defer lines
func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	testDefer()

	//defer execiton happens here (before the main func end '}')
}
func testDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//above function code stack works like
// main o/p-> hello testDefer() three two one
//testDefer o/p-> 4 3 2 1 0
