package main

import "fmt"

type Car struct {
	description string
	model       int
	make        string
}

func main() {

	fmt.Println("Demonstraiting Pointer to struct")

	car := Car{description: "S Class", model: 2023, make: "Merceides"}

	carPtr := &car

	fmt.Println("Details of car using pointer", carPtr)
	(*carPtr).description = "Hello"

	carPtr.model = 2022

	fmt.Println(car)

}
