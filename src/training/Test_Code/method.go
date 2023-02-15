package main

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) ToString() string {

	switch d {

	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

func main() {

	fmt.Println("Demo: Methods...")

	var direct Direction = North

	fmt.Println(direct, direct.ToString())

}
