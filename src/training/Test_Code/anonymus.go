package main

import (
	"fmt"
)

type myStruct struct {
	a, b int
	string
	float32
	float64
}

func main() {

	str := struct {
		name string
		pin  int
	}{name: "Sadanand D", pin: 560001}

	fmt.Println(str.pin)
	fmt.Println(str.name)

	var instance myStruct

	instance.int = 10
	instance.float32 = 121.2
	instance.float64 = 121.23233
	instance.string = "Mavenir..."

	fmt.Println(instance)
}
