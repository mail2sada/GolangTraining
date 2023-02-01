package main

import "fmt"

func main() {

	fmt.Println("Demonstraiting veriadic arguments...")

	veriadicStrings("Hello", "Mavenir")

	veriadicStrings("hello", "this", "is", "go", "lang", "training")

	res := veriadicConcat("hello", "this", "is", "go", "lang", "training")

	fmt.Println(res)

}

func veriadicStrings(ele ...string) {

	for _, b := range ele {
		fmt.Println(b)
	}
}

func veriadicConcat(ele ...string) string {
	var res string

	for _, b := range ele {
		res += b
		res += " "
	}
	fmt.Println(res)
	return res
}
