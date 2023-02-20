package main

import "fmt"

func main() {

	var a int = 30

	b := 30

	// if a > b {

	// 	fmt.Println("a is b")
	// } else {

	// 	fmt.Println("b is big")
	// }

	if a > b {
		fmt.Println("A is big")

	} else if a < b {
		fmt.Println("B is big")

	} else {
		fmt.Println("Both a and b are big")
	}

}
