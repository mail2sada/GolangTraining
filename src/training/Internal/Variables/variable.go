package main

import "fmt"

var a1 int

var b1 float64

func main() {

	fmt.Println("Demo: Variable declaration")

	var number int

	fmt.Println("Value of number is ", number)

	var testInt int = 10

	fmt.Println("Value os testInt is", testInt)

	var myIntr = 20

	fmt.Println("Value of myIntr is ", myIntr)

	fmt.Printf("Type of myInter is %T", myIntr)

	var a, b int = 10, 20

	fmt.Println("Value os a and b are", a, b)

	var x, y = 10, 22.33

	fmt.Printf("Type of x and y are %T %T", x, y)
}
