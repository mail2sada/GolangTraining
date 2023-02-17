package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo: Date and time...")

	t := time.Now()
	fmt.Println(t)

	expiry := t.Add(45 * time.Minute)

	//duration := time.Since(expiry)
	duration := time.Until(expiry)

	fmt.Println(duration)

}
