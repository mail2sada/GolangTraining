package main

import "fmt"

func main() {

	fmt.Println("Demo: Swtich case...")

	var a int = 100

	switch a {
	case 1:
		fmt.Println("Sunday..")
		fallthrough
	case 2:
		fmt.Println("Monday..")

	case 3:
		fmt.Println("Tuesday")

	case 7:
		fmt.Println("Sat")

	default:
		fmt.Println("Un handled case")
	}
}
