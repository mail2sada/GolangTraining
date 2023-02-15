package main

import "fmt"

var a = 100

func main() {

	fmt.Println("Demo: Variable declaration")

	var a int = 10

	var b, c string = "Hello  ", "World"

	var e = 20.99

	{
		var a = 500
		fmt.Println(a)
	}

	fmt.Println("Value of a is ", a)

	fmt.Println("Value of b and c are..", b, c)

	fmt.Println("Value of e is ", e)

	fmt.Printf("\n Type of e is %T", e)

}
