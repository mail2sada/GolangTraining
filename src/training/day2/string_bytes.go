package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	str := "Welcome to Mavenir"

	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {

		fmt.Printf("\nCharacter = %c Bytes = %v", str[c], str[c])
	}
}
