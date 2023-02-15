package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Demo: Strings....")

	var myStr string = "This is my \nfirst string.."

	var myNewString = `This is 
	Raw string using 
	backtick 
	\n "" !!@#''''`

	fmt.Println("Value os myStr is :", myStr)

	fmt.Println("Value os myNewString is:", myNewString)

	fmt.Println(myStr[3])

	fmt.Println(myStr[5])

	fmt.Println("Length of my string is ", len(myStr))
	fmt.Println("Length od my string using UTF8 package method", utf8.RuneCountInString(myStr))

	fmt.Printf("+++++++Address of myStr is %v", myStr)

	fmt.Println("String Concatination.....")

	myStr += myStr + "This test concatination..."
	fmt.Printf("=======Address of myStr is %v", myStr)

	

}
