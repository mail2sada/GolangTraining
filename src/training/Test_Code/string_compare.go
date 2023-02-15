package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Demo: String compare...")

	mystr := "Welcome to Mavenir..."

	mystr1 := "Hello World"

	mystr3 := "Welcome to Mavenir..."

	if mystr == mystr1 {

		fmt.Println("=========Both Strings are equal")
	} else {

		fmt.Println("=========Strings are not equal")
	}

	if mystr == mystr3 {
		fmt.Println("++++++++++Both Strings are equal")
	} else {
		fmt.Println("++++++++++Strings are not equal")

	}

	if strings.Compare(mystr, mystr1) == 0 {
		fmt.Println("<<<<<<<<<<<< String are equal...")
	} else {

		fmt.Println("<<<<<<<<<<<< Strings are not equal....")
	}

	if strings.Compare(mystr, mystr3) == 0 {
		fmt.Println("************* String are equal...")
	} else {

		fmt.Println("<************ Strings are not equal....")
	}


	
}
