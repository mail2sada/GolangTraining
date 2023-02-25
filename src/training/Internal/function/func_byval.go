package main

import "fmt"

func ByValue(x int) {
	x = 50
}

func ByRef(x *int) {
	*x = 100

}

func main() {
	fmt.Println("Demo Call by value...")

	a := 10

	ByValue(a)

	fmt.Println(a)

	ByRef(&a)

	fmt.Println(a)

}
