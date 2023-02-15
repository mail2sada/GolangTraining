package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = -1
	South           = iota
	East
	West
	northeast = 8
	southeast = iota
)

func main() {

	var d Direction = South

	fmt.Println("Direction is ", d)

}
