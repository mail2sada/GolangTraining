package main

import "fmt"

func printDetails(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Value is of type in ", a.(int))
	case string:
		fmt.Println("Value is of type string ", a.(string))
	case float32:
		fmt.Println("Value is of type float32 ", a.(float32))
	case float64:
		fmt.Println("Value is of type float64 ", a.(float64))
	default:
		fmt.Println("Value of type is not handled ", a)
	}
}

func main() {
	fmt.Println("Demo: empty interface")

	printDetails(10)

	printDetails(2.6)

	printDetails("Welcome to Mavenir Go lang training....")

	printDetails(1.970980392098312)
}
