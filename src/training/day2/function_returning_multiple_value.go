package main

import "fmt"

func main() {
	fmt.Println("Demonstration of function returning multiple values...")

	var a, b int = 10, 20

	fmt.Println("value of a and b before function call", a, b)

	a, b = swap(a, b)

	fmt.Println("value of a and b after 1 function call", a, b)

	_, b = swap(a, b)

	fmt.Println("value of a and b after 2 function call", a, b)

	a, _ = swap(a, b)
	fmt.Println("value of a and b after 3 function call", a, b)

}

func swap(a, b int) (int, int) {

	return b, a
}
