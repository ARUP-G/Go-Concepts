package main

import (
	"fmt"
)

func main() {

	fmt.Println("Struc module")
	// There is no inheritance in go
	// no super or parent , child
	Arup := User{"Arup", "@32fs", true, 23}
	fmt.Println(Arup)
	fmt.Printf("Details %+v \n", Arup) //%+v is used to display all the details
	fmt.Printf("Name is %v", Arup.Name)
}

type User struct { //fields name must start with capital to make them accessable/exportable
	Name   string
	Email  string
	Status bool
	Age    int
}
