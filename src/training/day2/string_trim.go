package main

import (
	"fmt"
	"strings"
)

//Please try it your self
// for your exploration TrimLeft(), TrimRight(), TrimSuffix(), TrimPrefix() and TrimSpace()

// Main method
func main() {

	// Creating and initializing string
	// Using shorthand declaration
	str1 := "!!Welcome to Mavenir !!"
	str2 := "@@This is the training material of Golang$$"

	// Displaying strings
	fmt.Println("Strings before trimming:")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2:", str2)

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.Trim(str1, "!")
	res2 := strings.Trim(str2, "@$")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}
