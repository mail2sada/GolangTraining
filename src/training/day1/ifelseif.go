package main

import "fmt"

func main() {

	// taking a local variable
	var v1 int = 700

	// checking the condition
	if v1 == 100 {

		// if condition is true then
		// display the following */
		fmt.Printf("Value of v1 is 100\n")

	} else if v1 == 200 {

		fmt.Printf("Value of a is 20\n")

	} else if v1 == 300 {

		fmt.Printf("Value of a is 300\n")

	} else {

		// if none of the conditions is true
		fmt.Printf("None of the values is matching\n")
	}
}
