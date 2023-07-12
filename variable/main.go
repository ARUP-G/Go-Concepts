package main

import "fmt"

var Name string = "Arup" // global | Publlic variable as var name starts with caps(N)
// Age := 22 // not allowed to use walrus operator as global declaration

func main() {
	fmt.Println("Types of variables")
	//variable defined without type
	// it is possible because as the variable is initiated go gan infer the variable type
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 // constant variable
	var remainingTickets = 50

	// variables can be declaread like (:= walrus operaator)
	newVar := "String"
	fmt.Printf("Type of newVar is: %T \n", newVar)

	// variable defined with type
	var hello string = " Hi I am a string type variable"
	fmt.Println(hello)

	fmt.Printf("Types of variables are: %T , %T \n", conferenceName, conferenceTickets)

	fmt.Printf("Variable type %T \n", Name) // global access

	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

}
