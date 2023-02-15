package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Demo: String split....")

	myString := "Welcome to mavenir Golang training"

	strSlice := strings.Split(myString, " ")

	for idx, val := range strSlice {

		fmt.Println(" value at idx", val, idx)

	}

}
