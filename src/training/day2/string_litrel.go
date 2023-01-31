package main

import "fmt"

func main() {

	// Creating and initializing a
	// variable with a string literal
	// Using double-quote
	My_value_1 := "Welcome to Mavenir"

	// Adding escape character
	My_value_2 := "Welcome!\nMavenir"

	// Using backticks
	My_value_3 := `Hello!Mavenir`

	// Adding escape character
	// in raw literals
	My_value_4 := `Hello!\nMavenir`

	// Displaying strings
	fmt.Println("String 1: ", My_value_1)
	fmt.Println("String 2: ", My_value_2)
	fmt.Println("String 3: ", My_value_3)
	fmt.Println("String 4: ", My_value_4)
}
