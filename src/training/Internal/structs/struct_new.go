package main

import "fmt"

type Strudent struct {
	roll_no int
	name    string
}

func main() {
	fmt.Println("Demo: struct_new")

	a := new(Strudent) //heap // malloc //new java and c++

	a.roll_no = 01
	a.name = "Hello"

	fmt.Printf("Type of a is %T", a)

	b := Strudent{}

	fmt.Println(b)

}
