package main

import "fmt"

func main() {
	// airthmatic operators
	// +, -, *, / and %

	var a, b int = 20, 10

	c := a + b
	fmt.Println("Value of (a+b) is ", c)

	d := a - b
	fmt.Println("Value of (a-b) is ", d)

	e := a * b
	fmt.Println("Value of (a*b) is ", e)

	f := a / b
	fmt.Println("Value of (a/b) is ", f)

	g := a % b
	fmt.Println("Value of (a%b) is ", g)

	// relational operators
	// <, >, ==, !=, <=, >=

	//a, b := 10, 20

	res := (a < b)
	fmt.Println(res)

	res1 := (a > b)
	fmt.Println(res1)

	res2 := (a == b)
	fmt.Println(res2)

	res3 := (a != b)
	fmt.Println(res3)

	res4 := (a <= b)
	fmt.Println(res4)

	res5 := (a >= b)
	fmt.Println(res5)

	// Logical operator
	// Logical &&, logical || and logical !

}
