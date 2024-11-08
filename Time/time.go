package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Lets play with the time library !")

	presentTime := time.Now()

	fmt.Println(presentTime)

	// Lets play with the time format 
	// This is the fix format and constant "01-02-2006 15:04:05 Monday"

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}