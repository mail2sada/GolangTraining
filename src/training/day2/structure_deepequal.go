package main

import (
	"fmt"
	"reflect"
)

// Creating a structure
type Author struct {
	name     string
	branch   string
	language string
}

// Main function
func main() {

	// Creating variables
	// of Author structure
	a1 := Author{
		name:     "Soana",
		branch:   "CSE",
		language: "Perl",
	}

	a2 := Author{
		name:     "Soana",
		branch:   "CSE",
		language: "Perl",
	}

	a3 := Author{
		name:     "Dia",
		branch:   "CSE",
		language: "Perl",
	}

	// Comparing a1 with a2
	// Using DeepEqual() method
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))

	// Comparing a2 with a3
	// Using DeepEqual() method
	fmt.Println("Is a2 equal to a3: ", reflect.DeepEqual(a2, a3))
}
