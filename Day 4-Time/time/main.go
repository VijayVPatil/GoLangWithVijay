package main

import (
	"fmt"
	"time"
)

func main() {
	//Time package

	presentTime := time.Now()
	fmt.Println(presentTime)

	//Now to print in follwoing format we need to pass the below specific values to get the format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
