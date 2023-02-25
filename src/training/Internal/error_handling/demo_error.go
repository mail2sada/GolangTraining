package main

import "fmt"

func Divide(a, b int) int {

	if b == 0 {
		panic("Divide by zero error") // to raise exception
	}

	c := a / b
	return c

}

func main() {

	defer func() {
		a := recover() // handle exception
		if a != nil {
			fmt.Println("Recovering from panic", a)
		}

	}()

	fmt.Println("Demo: Error handling...")

	fmt.Println("Divide ", Divide(10, 20))

	fmt.Println("Divide ", Divide(30, 20))

	fmt.Println("Divide ", Divide(10, 0))

	fmt.Println("Divide", Divide(55, 100))

}
