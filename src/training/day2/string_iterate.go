package main

import "fmt"

// Main function
func main() {

	// String as a range in the for loop
	for index, s := range "Hello Mavenir, Great to have you here!!" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}

	for index, s := range `Hello Mavenir, \n \r Great to have you here!!` {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}
