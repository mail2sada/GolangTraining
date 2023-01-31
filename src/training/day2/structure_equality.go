package main

import "fmt"

// Creating a structure
type Author struct {
	name      string
	branch    string
	language  string
	Particles int
}

// Main function
func main() {

	// Creating variables
	// of Author structure
	a1 := Author{
		name:      "Moana",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	a2 := Author{
		name:      "Moana",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	a3 := Author{
		name:      "Dona",
		branch:    "CSE",
		language:  "Python",
		Particles: 38,
	}

	// Checking if a1 is equal
	// to a2 or not
	// Using == operator
	if a1 == a2 {

		fmt.Println("Variable a1 is equal to variable a2")

	} else {

		fmt.Println("Variable a1 is not equal to variable a2")
	}

	// Checking if a1 is equal
	// to a2 or not
	// Using == operator
	if a2 == a3 {

		fmt.Println("Variable a2 is equal to variable a3")

	} else {

		fmt.Println("Variable a2 is not equal to variable a3")
	}
}
