package main

import "fmt"

func main() {
	fmt.Println("Demo: Using errorf")
	sampleErr := fmt.Errorf("Err is: %s", "database connection issue")
	fmt.Println(sampleErr)
}
