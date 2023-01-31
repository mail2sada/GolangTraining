package main

import (
	"fmt"
	"strings"
)

func main() {

	var i int = 4
	switch i {
	case 1:
		fmt.Println("Its One")
	case 2:
		fmt.Println("Its Two")
	case 3:
		fmt.Println("Its Three")
		fallthrough
	case 4:
		fmt.Println("\"Its Four\"")
		fallthrough
	case 5:
		fmt.Println("Its Five")
		fallthrough
	default:
		fmt.Println("Sorry, I did not handle you....")
	}

	var str string = "NAMASTE"

	switch strings.ToLower(str) {
	case "hello":
		fmt.Println("Greeting in english")
	case "bonjour":
		fmt.Println("Greeting in french")
	case "namaste":
		fmt.Println("Greeting in Hindi")
	}

}
