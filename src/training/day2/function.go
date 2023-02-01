package main

import "fmt"

func main() {
	fmt.Println("Demposntraining of funtion")

	firstFuntion()
	myFuncWithArgument("Passing Arguments...")

	res := add(10, 20)

	fmt.Println("result is ", res)

}

func firstFuntion() {

	fmt.Println("This is our first function...")
}

func myFuncWithArgument(str string) {

	fmt.Println("This the string that we have passed...", str)
}

func add(a int, b int) int {

	fmt.Println("Demonstringing returning values...")

	c := a + b

	return c
}
