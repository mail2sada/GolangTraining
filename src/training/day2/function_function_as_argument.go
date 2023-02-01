package main

import "fmt"

func main() {
	fmt.Println("Demonstraiting hire order function")
	higherOrderFunction(func(str string) {
		fmt.Println(str)
	}, "This is higher order function demonstration...")

	fn := demonstrationHigerOrder()

	fn("Hello world")

}

func higherOrderFunction(fn func(str string), str string) {
	fn(str)
}

func demonstrationHigerOrder() func(str string) {
	return (func(str string) {
		fmt.Println("Value of str is str", str)
	})
}
