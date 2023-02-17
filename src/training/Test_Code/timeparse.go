package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo time parse...")

	t, err := time.Parse("2006-01-02 03:04:05", "2023-02-16 11:49:21")

	t1, err := time.Parse("06/Jan/02 03-04-05 MST", "23/Feb/18 09-49-21 IST")

	fmt.Printf("\nTime is %d:%d:%d\n", t.Day(), t.Hour(), t.Minute())

	fmt.Println(t1)

	fmt.Println(t, err)

	timeStr := t1.Format("2006/January/_2 03:04:05 MST")

	fmt.Println(timeStr)

	x := t.Unix()

	fmt.Println("Unix time stamp is ", x)

	t3 := time.Unix(x, 0)

	fmt.Println(t3)
}
