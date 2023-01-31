package main

import "fmt"

func main() {

	// Creating and initializing strings
	str1 := "Welcome"
	str2 := "Mavenir"

	// Using += operator
	str1 += str2
	fmt.Println("String: ", str1)

	str1 += "This is the tutorial of Go language"
	fmt.Println("String: ", str1)

	str2 += " Golang training"
	fmt.Println("String: ", str2)

}
