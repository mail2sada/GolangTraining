package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "hello"

	str = strings.ToUpper(str)

	var str1 string = "Welcome to \tMavenir training"

	var str2 string = `Hello hopw are you "" @##$34-30943
	sdslkjdsdj
	skdjslkjdls`

	fmt.Println(str1)

	fmt.Println(str2)
}
