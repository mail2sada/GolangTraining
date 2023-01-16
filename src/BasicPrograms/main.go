package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("[MAIN]")
	log.SetFlags(log.Flags())
}

func main() {
	var a, b = 10, false
	var x int
	x = 0
	var y string
	y = ""
	fmt.Println(x, y)
	fmt.Printf("Tye of a is %T and type of b is %T", a, b)
	fmt.Println("Welcome To Go Lang Basic Programs")
	testYourCode()
}

func testYourCode() {
	a := 10
	fmt.Println("Value of a is ", a)
	log.Fatal("What is this...?")
}
