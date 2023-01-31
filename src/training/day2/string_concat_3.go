package main

import "fmt"

func main() {

	// Creating and initializing strings
	str1 := "Mavenir"
	str2 := "Go"
	str3 := "Lang"
	str4 := "training"

	// Concatenating strings using
	// Sprintf() function
	result := fmt.Sprintf("%s %s %s %s", str1,
		str2, str3, str4)

	fmt.Println(result)
}
