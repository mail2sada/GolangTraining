package main

import (
	"fmt"
	"time"
)

func main() {

	var t time.Time = time.Now()

	fmt.Println("Current time is ", t)

	fmt.Println(t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

	if tm := time.Now(); tm.Month() == time.February && tm.Day() == 14 {

		fmt.Println("Current year is ", tm.Year())
		fmt.Println("This month is Feb.. and day is 13")
	}

}
