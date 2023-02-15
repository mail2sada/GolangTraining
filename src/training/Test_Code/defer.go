package main

import "fmt"

func main() {

	fmt.Println("Demo: Defer...")

	x := 20

	defer func(a int) {
		fmt.Println(a)
	}(x)

	defer myTest()

	x = 50

	fmt.Println("After defer...")

}

func myTest() {

	fmt.Println("In myTest")

	defer fmt.Println("This is myTest()")

}
