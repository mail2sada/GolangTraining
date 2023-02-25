package main

import "fmt"

func main() {

	fmt.Println("Demo: Anonymous struct")

	a := struct {
		name string
		msg  string
	}{name: "123", msg: "how are you"}

	fmt.Println(a)
}
