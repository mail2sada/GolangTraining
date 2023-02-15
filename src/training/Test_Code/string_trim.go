package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Demo: String trim...")

	myString := `     this is go lang training 
	in mavenir..,    `

	fmt.Println("Before trim ...", myString)
	myString = strings.Trim(myString, ` `)

	fmt.Println("After trim...", myString)

	

}
