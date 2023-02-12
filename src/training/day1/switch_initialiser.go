package main

import "fmt"
import "time"

func main() {

	fmt.Println("Demo: Switch Initialiser...")


	switch today := time.Now(); {

	case today.Day()<=7:
		fmt.Println("First Week of the month")
	case today.Day()<= 14:
		fmt.Println("Second Week of the Month ")
	case today.Day() <= 21:
		fmt.Println("Third week of the month..")
	case today.Day() <= 28:
		fmt.Println("Fourth week of the month...")
	default:
		fmt.Println("Last week of the month...")
	}
}