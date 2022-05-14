package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// Formatting the time and date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// Note: this is only format need to stick with it

	createdDate := time.Date(2020, time.March, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
