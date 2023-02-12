package main

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)



func (d Direction) ToString()string{
	res := ""
	switch d {
		case North:
			res = "North"
		case South:
			res = "South"
		case East:
			res = "East"
		case West:
			res = "West"
	}
	return res
}

func main() {

	fmt.Printf("Demo: Simulationing enum",a)

	d := North

	var direct Direction = South

	const fixedDirect Direction = East

	const oppFixedDirect Direction = West

	fmt.Printf("%s\n%s\n%s\n%s\n", d.ToString(), direct.ToString(), fixedDirect.ToString(), oppFixedDirect.ToString)
}