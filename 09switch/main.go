package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Ludo game")

	var count int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		diceNumber := rand.Intn(6) + 1 // number range is always non-inclusive so 1 is added
		fmt.Println("Dice numbre is: ", diceNumber)

		if diceNumber == 1 && count == 0 {
			fmt.Println("Dice number is 1, you can open or go 1 unit")
			count = -1
		} else {
			//go switch has no autometic fallthrough
			switch diceNumber {
			case 1:
				fmt.Println("Dice number is 1, you can go 1 unit")
			case 2:
				fmt.Println("You got 2, move 2 unit")
			case 3:
				fmt.Println("You got 3, move 3 unit")
			case 4:
				fmt.Println("You got 4, move 4 unit")
			case 5:
				fmt.Println("You got 5, move 5 unit + you will get 6 bonus move")
				fallthrough //this autometically calls the next case
			case 6:
				fmt.Println("You got 6, move 6 unit")
			default:
				fmt.Println("Roll again")
			}
		}
	}

}
