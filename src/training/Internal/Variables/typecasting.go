package main

import "fmt"

func main() {

	fmt.Println("Demo: Type Casting..")

	var a float64 = 10

	var b int = 20

	a = float64(b)

	fmt.Println("Value of a , b are ", a, b)

}
