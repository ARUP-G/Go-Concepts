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
	fmt.Printf("%s \n", finalJson) // output as string
}
