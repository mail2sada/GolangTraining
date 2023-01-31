package main

import "fmt"

var global int = 100

func main() {
	var local int = 200

	fmt.Println(global)
	fmt.Println(local)

	{
		var block int = 300
		fmt.Println(block)
	}
	//block is not accessible outside the block
	//fmt.Println(block)
}
