package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time")

	presentTime := time.Now()
	fmt.Println("Time of now", presentTime)

	// format is "01-02-2006 15:04:05 Monday" always
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	newDate := time.Date(2020, time.November, 11, 21, 0, 0, 0, time.UTC)
	fmt.Println("", newDate)
	fmt.Println(newDate.Format("01-02-2020 Monday"))
	//fmt.Println(newDate)

	// to build for mac and linux
	// command for linux file-> GOOS="linux" go build
	// command for mac file -> GOOS="darwin" go build

}
