package main

import "fmt"

func main() {
	fmt.Println("Hello world from Golang...")
}

func MyTestPrint(elements ...any) {

	for _, val := range elements {

		fmt.Println(val)
	}
}
