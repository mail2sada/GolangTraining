package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Wonderful....")

	fmt.Println("Current time is ", time.Now())
	fmt.Println("Today is ", time.Now().Day())
	fmt.Println("Today is ", time.Now().Weekday())
	fmt.Println("Today is in ", time.Now().Month())
	fmt.Println("Today is in year ", time.Now().Year())
}
