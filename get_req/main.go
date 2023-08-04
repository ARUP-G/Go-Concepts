package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Get request from server")
	PerformGetRequest()
}

func PerformGetRequest() {

	const myUrl = "http://localhost:8080/form.html"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// status code
	fmt.Println("Status Code", response.StatusCode)

	//Content length
	fmt.Println("Content length: ", response.ContentLength)

	// Read all the content
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	//content display
	// Content is in byte format, so wrap it in string
	fmt.Println("Content of the Url body:", string(content))

	// Another way to read the content of the page

	var responseString strings.Builder
	content2, _ := io.ReadAll(response.Body)
	//String the response in String Builder and holding the data by write method
	buteCount, _ := responseString.Write(content2)

	// Same as Content count
	fmt.Println("Byte count : ", buteCount)

	// Getting the data as a srting format from the string bulder
	fmt.Println("Content of the page:", responseString.String())

	if err != nil {
		panic(err)
	}

}
