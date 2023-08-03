package main

import (
	"fmt"
	"net/url"
)

// '&' in url works as comma
// change the "link" with a actual link
const myUrl string = "link"

func main() {
	fmt.Println("URL Handeling")

	fmt.Println(myUrl)

	//parsing -> information gathering
	result, err := url.Parse(myUrl)
	checkNilErr(err)

	fmt.Println("Result of scheme:", result.Scheme)
	fmt.Println("Result of host:", result.Host)
	fmt.Println("Result of path:", result.Path)

	fmt.Println("Result of port:", result.Port())
	fmt.Println("Result of raw-query:", result.RawQuery)

	//Query parameter
	qparam := result.Query()
	fmt.Printf("Type of qparam: %T \n", qparam) //-> url.values (key value parameter)

	fmt.Println("from url: ", qparam["index"])
	fmt.Println("from url: ", qparam["list"])

	//as qparam are key value pair (map)
	for index, val := range qparam {
		fmt.Printf("Parameter %v is:%v \n", index, val)
	}

	partsOfUrl := &url.URL{ // have to be this way & is mandotory

		Scheme: "https",
		Host:   "www.youtube.com",
		Path:   "/watch",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

// error checker
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
