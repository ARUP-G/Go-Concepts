package main

import (
	"fmt"
)

func main() {

	fmt.Println("Method module")
	// There is no inheritance in go
	// no super or parent , child
	Arup := User{"Arup", "@32fs", true, 23}
	fmt.Println(Arup)
	fmt.Printf("Details %+v \n", Arup) //%+v is used to display all the details
	fmt.Printf("Email is %v\n", Arup.Email)
	Arup.getStatus()
	Arup.newMail() //creates a copy of the user
	fmt.Printf("Email is %v\n", Arup.Email)

}

type User struct {
	//fields name must start with capital to make them accessable/exportable
	Name   string
	Email  string
	Status bool
	Age    int
}

// method
func (u User) getStatus() {
	fmt.Println("Online:", u.Status)
}
func (u User) newMail() {
	u.Email = "@new"
	fmt.Println("New email is: ", u.Email)

}
