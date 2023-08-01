package main

// code for write and read a file
import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File access module")

	// it is gooing to be written in the file
	content := "This will be added to file - File access with golang"

	//os package is used to creat the file
	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err) // error check

	//to write the content
	length, err := io.WriteString(file, content)

	checkNilErr(err) //error check

	//length of the content
	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {

	// go program reads the data of the file as data byte
	databyte, err := os.ReadFile(filename)

	checkNilErr(err)

	//wrap databyte to string to make it readable for human
	fmt.Println("Contents of the file:", string(databyte))
}

// error handeling function
func checkNilErr(err error) {
	if err != nil {
		//panic(err) stop pprogram and show the error
		panic(err) // panic keyword will stop the program
	}

}
