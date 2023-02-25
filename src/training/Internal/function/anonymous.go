package main

import "fmt"

func main() {

	fmt.Println("Demo: Anonymous functions...")

	func() {
		fmt.Println("Hello..")
	}()

	func(str string) {
		fmt.Println(str)
	}("Hello")

	myStr := func(str string) string {

		return str + " from Anonymous"

	}("Hello..")

	fmt.Println(myStr)

	fmt.Println(func(str string) string {

		return str + " from Anonymous"

	}("Hello.."))

	myFunc := func(test string) {
		fmt.Println("Wanted  ", test)
	}

	myFunc("test")

	myFunc("Best")
}
