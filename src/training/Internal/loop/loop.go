package main

import "fmt"

func main() {

	fmt.Println("Demo loop")

	fmt.Println("======= Simple infinte loop")

	a := 1

	for a <= 10 {
		fmt.Println("Hello.....")
		a++
	}

	for a = 50; a < 100; a++ {

		fmt.Println("Welcome...")
	}

	for i := 100; i < 400; i++ {
		fmt.Println("Value of i is ", i)
	}
}
