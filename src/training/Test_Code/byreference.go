package main

import "fmt"

// func swap(a, b int) {

// 	c := a
// 	a = b
// 	b = c
// }

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {

	fmt.Println("Demo")

	var a, b int = 10, 20

	swap(&a, &b)

	fmt.Println(a, b)
}

func returnPtr() *int {
	a := 10

	return &a
}
