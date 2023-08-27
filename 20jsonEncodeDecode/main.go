package main

import (
	"encoding/json"
	"fmt"
)

// make a struct
type account struct {
	Name     string   `json:"accholder_name"`
	AcNum    int      `json:"accnumber"`
	Branch   string   `json:"branch"`
	Password string   `json:"-"`
	AcType   []string `json:"acctype"`
}

func main() {
	fmt.Println("Json Encodeing & Decoading")
}

func Encodeing() {
	// slice of struct
	SBI := []account{
		{"Arup", 74589, "RPG", "sef41", []string{"Savings", "Current"}},
		{"Ron", 848826, "Kol", "dfe816", []string{"Saving", "Current", "Treadnig"}},
		{"Dev", 18466, "KIS", "adsw88", nil},
	}

	//encoding
	finalJson, err := json.MarshalIndent(SBI, "", "\t")
	if err != nil {
		panic(err)
	}
	// output as string
	fmt.Printf("%s \n", finalJson)
}

func Decoading() {

	jsonData := []byte(`
	{
		"holdername": "Ron",
		"accountNum": 848826,
		"branch": "Kol",
		"_": "dfe816",
		"acctype": ["Saving","Current","Treadnig"]
	}
	`)

	// create a var of type struct to store the json data
	var accDetails account

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json Data is Valid")
		// unmarshaling the jsonData and saving to accDetails
		json.Unmarshal(jsonData, &accDetails)
		fmt.Printf("%#v\n", accDetails)
	} else {
		fmt.Println("Json Not Valid !!")
	}
}
