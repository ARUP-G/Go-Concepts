package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("Web request module")

	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Type of responce %T \n", response)
	//fmt.Println("Responded", response)

	// ReadResponse nor Response.Write ever closes a connection.
	defer response.Body.Close() // callers responsibility to close the connnection

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Println("Response body is: ", content)
}
func checkNilErr(err error) {
	if err != nil {
		//panic(err) stop pprogram and show the error
		panic(err) // panic keyword will stop the program
	}

}
