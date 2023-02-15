package main

import "fmt"

func swap(a, b int) (int, int) {

	return b, a
}

func area(a, b int) (areaSquare int, areaRect int) {

	areaSquare = a * a
	areaRect = a * b
	return
}

func main() {

	fmt.Println("Demo: Functions")

	var a, b int = 10, 20

	a, b = swap(a, b)

	fmt.Println("Value of a and b are ", a, b)

	var x, y int = 10, 11

	areaSqure, _ := area(x, y)

	fmt.Println("Area of square is ", areaSqure)

	//fmt.Println("Area of rect is ", areaRect)

}
