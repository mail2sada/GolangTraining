package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: demonstration of date ")

	var t1 time.Time = time.Now()

	fmt.Println("Current time is:", t1)

	var t2 time.Time = time.Date(2023, time.February, 28, 11, 53, 59, 0, new(time.Location))

	fmt.Println("formed date value is ", t2)

	var duration time.Duration = 1 * 60 * 60 * 1000 * 1000 * 1000

	fmt.Println("duration is ", duration)

	//Adding and subtracting time

	tnew := t1.Add(time.Hour * 1)

	duration = tnew.Sub(t1)

	fmt.Println("duration ", duration)

}
