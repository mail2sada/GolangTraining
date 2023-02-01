package main

import "fmt"

func main() {
	fmt.Println("Demonstraiting Slice as argument to function")

	elements := []string{"Go_", "lang_", "training"}

	ele := concat_string(elements)
	fmt.Println("This is my concatinated string ", ele)
}

func concat_string(elements []string) (res string) {

	for _, ele := range elements {
		res += ele
	}
	return
}
