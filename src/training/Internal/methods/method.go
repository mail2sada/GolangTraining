package main

import "fmt"

type Direction uint8

const (
	North Direction = iota
	South
	East
	West
)

func (a Direction) ToString() string {

	switch a {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	}
	return ""

}

func main() {

	fmt.Println("Demo Methods")

	var dirct Direction = North

	fmt.Println(dirct.ToString())
}
