package main

import (
	"fmt"
	"unicode/utf8"
)

// Main function
func main() {

	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to Mavenir ??????"

	// Finding the length of the string
	// Using len() function
	length1 := len(mystr)

	// Using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystr)

	// Displaying the length of the string
	fmt.Println("string:", mystr)
	fmt.Println("Length 1:", length1)
	fmt.Println("Length 2:", length2)

}
