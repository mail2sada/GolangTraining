package main

import "fmt"

type Address struct {
	building string
	city     string
	pin      int
}

type Prompted struct {
	string
	int
	float32
	float64
	uint16
	Address
}

func main() {
	fmt.Println("Demo: Prompted fields")

	var prompted Prompted = Prompted{int: -123, string: "Hello", float32: 10.55, float64: 10.99999, uint16: 1234}

	fmt.Println(prompted)
}
