package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Give me a number")

	// for input
	reader := bufio.NewReader(os.Stdin)
	// taking input till end line with error handeling
	input, _ := reader.ReadString('\n') //input is type of string

	fmt.Println("Your number is: ", input)

	// addOne := input + 1 // error as  input is of string type
	addOne, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strings.TrimSpace clears the '\n' error (parsing "5\r\n": invalid syntax)

	if err != nil {
		fmt.Println(err)
		//panic(err) // this will end the program

	} else {
		fmt.Println("Added 1 to the number ", addOne+1)
	}

}
