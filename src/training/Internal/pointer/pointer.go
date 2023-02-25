package main

import "fmt"

func main() {

	var str string = "Hello world"

	fmt.Println("Value of str is", str)

	var ptrStr *string = &str

	fmt.Println(ptrStr)

	fmt.Println(&str)

	fmt.Println(*ptrStr)

	var ptrPtrStr **string = &ptrStr

	fmt.Println(**ptrPtrStr) //contents of str

	fmt.Println(*ptrPtrStr) //address of str

	fmt.Println(ptrPtrStr) // addresss of ptrstr
}
