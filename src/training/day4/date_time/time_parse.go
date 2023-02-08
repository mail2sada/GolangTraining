package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo: Time parse...")

	t, _ := time.Parse("2006 01 02 15 04", "2023 02 16 16 50")
	fmt.Println(t.YearDay()) // 315
	fmt.Println(t.Weekday()) // Wednesday

	t, _ = time.Parse("2006 01 02 15 04", "2023 02 13 0 00")
	fmt.Println(t.YearDay())
	fmt.Println(t.Weekday())
}
