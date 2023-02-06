package main

import "fmt"

func MySPrint(elements ...interface{}) string {

	var res = ""

	for _, val := range elements {
		res += fmt.Sprintf("%v", val)
	}

	return res

}

func main() {
	fmt.Println("Demo: Veriadic functions with Interface")

	fmt.Println(MySPrint("Hello ", 10, 22.64))
}
