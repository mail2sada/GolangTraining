package main

import "fmt"

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	// Declaring and initializing a
	// struct using a struct literal
	var a1 Address = Address{"Akshay", "Bangalore", 560045}

	fmt.Println("Address1: ", a1)

	// Naming fields while
	// initializing a struct
	var a2 = Address{Name: "Anikaa", city: "Chennai",
		Pincode: 440001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Delhi", city: "Bangalore", Pincode: 560102}
	fmt.Println("Address3: ", a3)
}
