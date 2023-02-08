package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("Berlin Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo Time: %s\n", now.In(loc))
}
