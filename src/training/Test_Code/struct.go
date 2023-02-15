package main

import (
	"fmt"
	"reflect"
)

// type address struct {
// 	name    string
// 	apt     string
// 	street  string
// 	city    string
// 	pincode int
// }

type address struct {
	name, apt, street, city string
	pincode                 int
}

func main() {

	fmt.Println("Demo: Struct...")

	var myAddress address = address{name: "Sadanand", apt: "Sobha Daffodils", street: "HSR", city: "Bangalore", pincode: 560102}

	var yourAddress = address{name: "Sadanand", apt: "Sobha Daffodils", street: "HSR", city: "Bangalore", pincode: 560102}

	fmt.Println(myAddress.name)
	fmt.Println(myAddress.apt)
	fmt.Println(myAddress.street)
	fmt.Println(myAddress.city)
	fmt.Println(myAddress.pincode)

	fmt.Println(yourAddress)

	yourAddress.apt = "Poorvi Lotus"
	yourAddress.street = "BTM"

	fmt.Println(yourAddress)

	var testAdrress = new(address)

	testAdrress.name = "Mavenir"
	testAdrress.apt = "MFAR"
	testAdrress.street = "Manyata"
	testAdrress.city = "Bangalore"
	testAdrress.pincode = 560045

	fmt.Println(*testAdrress)

	if myAddress == yourAddress {
		fmt.Println("Equals...")
	} else {
		fmt.Println("Not equals...")
	}

	if reflect.DeepEqual(yourAddress, myAddress) {
		fmt.Println("Equals...")
	} else {
		fmt.Println("Not equals...")
	}

}
