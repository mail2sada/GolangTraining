package main

import "fmt"

// Main function
func main() {

	// Creating and initializing strings
	// using shorthand declaration
	str1 := "Mavenir"
	str2 := "Mac"
	str3 := "MavenirMav"
	str4 := "Mavenir"

	// Checking the string are equal
	// or not using == operator
	result1 := str1 == str2
	result2 := str2 == str3
	result3 := str3 == str4
	result4 := str1 == str4

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)

	// Checking the string are not equal
	// using != operator
	result5 := str1 != str2
	result6 := str2 != str3
	result7 := str3 != str4
	result8 := str1 != str4

	fmt.Println("\nResult 5: ", result5)
	fmt.Println("Result 6: ", result6)
	fmt.Println("Result 7: ", result7)
	fmt.Println("Result 8: ", result8)

}
