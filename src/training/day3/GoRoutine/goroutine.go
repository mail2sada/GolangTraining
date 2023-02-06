package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Go routine...")

	go fmt.Println("Welcome")

	go fmt.Println(" Go Lang training")

	time.Sleep(1 * time.Second)
}
