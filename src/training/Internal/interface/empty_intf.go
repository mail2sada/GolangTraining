package main

import "fmt"

type st struct {
}

func MyCustomrPrintln(elements ...interface{}) {

	for _, val := range elements {
		fmt.Print(val)
	}
	fmt.Println()

}

func Add(a, b interface{}) interface{} {

	switch a.(type) {
	case int:

		aInt := a.(int)
		bInt := b.(int)

		return aInt + bInt

	case float64:
		afloat := a.(float64)
		bfloat := a.(float64)

		return afloat + bfloat
	case string:

		astring := a.(string)
		bstring := b.(string)

		return astring + bstring

	}
	return 0
}

func main() {

	a := Add(10, "Test")

	b := Add(22.55, 13.14)

	c := Add("Hello..", "Friends")

	MyCustomrPrintln(a, b, c)

	MyCustomrPrintln("Hello ", 123, "abcde", 22.55)

}
