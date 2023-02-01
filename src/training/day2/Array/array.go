package main

import "fmt"

func main() {

	// Shorthand declaration of array
	arr := [4]string{"hello", "mav", "Mavenir", "Go Lang training..."}

	// Accessing the elements of
	// the array Using for loop
	fmt.Println("Elements of the array:")

	for i := 0; i < 4; i++ {
		fmt.Println(arr[i])
	}
}
