package main

import (
	"encoding/json" //for json
	"fmt"
)

func main() {
	Arup := DevOpsEng{8628, "Arup Das", 529800.22}

	//fmt.Printf("Delatis %+v", Arup)
	byteOut, _ := json.Marshal(Arup)
	//fmt.Println("Json marsal: ", string(byteOut))

	var a DevOpsEng

	json.Unmarshal(byteOut, &a) //byteOut get stores into a

	fmt.Println("a", a)
}

type DevOpsEng struct {
	EmpID int
	FName string
	Sal   float32
}
