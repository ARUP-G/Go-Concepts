package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number")

	// comma ok / comma error syntax
	//input is try & _ is catch as go has no error handeling
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating is:", input)
	fmt.Printf("Type of the input is %T", input)

}
