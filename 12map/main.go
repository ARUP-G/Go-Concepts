package main

import "fmt"

// its a  key, value pair
func main() {
	fmt.Println("Maps")
	//map[key]value
	languages := make(map[string]string)

	//adding
	languages["JS"] = "JavaScrip"
	languages["PY"] = "Python"
	languages["Go"] = "Golang"
	languages["MD"] = "Markdown"

	fmt.Println("Map: ", languages)
	fmt.Println("JS is short form of:", languages["JS"])

	//delete
	delete(languages, "MD")
	fmt.Println("After deleting:", languages)

	//loop

	for key, value := range languages {
		fmt.Printf("For key %v, valuse is %v \n", key, value)
	}

	//using comma ok syntax
	fmt.Println("Other loop")
	for _, value := range languages {
		fmt.Printf("For key , valuse is %v \n", value)
	}
}
