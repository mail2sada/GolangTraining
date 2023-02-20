package main

import "fmt"

const a1, b1 = 20, "Hello"

func main() {

	fmt.Println("Demo: Constants ")

	const a = 10

	const test int = 10

	const mycnst uint8 = 10

	fmt.Println("Valus of a is ", a)

	fmt.Println("Value of test is", test)

	fmt.Println("Value of mycnst is ", mycnst)

	fmt.Println(a1, b1)
}
