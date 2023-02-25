package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Demo: String....")

	var str string = "This is my test string..\n You must be ssing this in newline as we Added newline,"

	var testStr string = `\n "" !@#$$%%^'' ;';';';';'
	askjdhsjkadhk kjhk jhkjdas
	shdkjsahdjksahd ksa
	
	oqwhlkjdhlshkj
	oeyqwueyiquw
	
	asjdhkjsahd`

	x := len(testStr)
	y := utf8.RuneCountInString(testStr)

	//range, len

	fmt.Println(str)

	fmt.Println(testStr)
}
