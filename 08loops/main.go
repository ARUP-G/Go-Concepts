package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lopps module")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saterday", "Sunday"}

	//fmt.Println("Days:", days)

	fmt.Println("Length of the slice:", len(days))
	fmt.Println("7:", days[6])

	//normal
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// using range
	// for i := range days {
	// 	fmt.Println("Days are:", days[i])
	// }

	//for each
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and day is %v \n", index, day)
	// }

	//using comma ok
	// for _, day := range days {
	// 	fmt.Printf("Day is: %v\n", day)
	// }

	anyValue := 1

	//like while loop
	for anyValue < 10 {

		// if anyValue == 3 {
		// 	break // stops at the value
		// }
		if anyValue == 1 {
			goto lco
		}
		if anyValue == 7 {
			anyValue++ //if not incremented the loopo will hit infinity
			continue   // skips the value
		}
		fmt.Println("Value is: ", anyValue)
		anyValue++
	}

lco:
	fmt.Println(" Hit no 1")
}
