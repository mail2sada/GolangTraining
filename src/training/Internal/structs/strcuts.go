package main

import "fmt"

type address struct {
	name     string
	building string
	street   string
	pincode  int
}

func main() {

	fmt.Println("Demo Package")

	var addr address = address{name: "Yogesh P", building: "Manyata", street: "Nagawara", pincode: 560040}

	fmt.Println(addr.building)

	fmt.Println(addr.name)

	fmt.Println(addr.pincode)

	fmt.Println(addr.street)

	fmt.Println(addr)

	addr.name = "Guruprasad G"

	addr.street = "Sahakar nagar"

	name := addr.name

	fmt.Println(name)

	fmt.Println(addr)
}
