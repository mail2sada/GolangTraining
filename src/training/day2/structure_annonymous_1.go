package main

import "fmt"

func main() {

	fmt.Println("Demonstraining Annonymous structre")

	ele := struct {
		name    string
		address string
	}{
		name:    "Vijay",
		address: "Bangalore 560111",
	}

	fmt.Println(ele)

}
