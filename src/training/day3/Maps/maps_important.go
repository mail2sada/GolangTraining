package main

import "fmt"

func main() {

	//this is empty map
	var details = map[string]int{}
	fmt.Println("Printing map details", details)

	//Add some values

	details["Mark"] = 55
	fmt.Println("Printing map after update ", details)

	//creating map using make function

	var emptyMap = make(map[string]string)

	fmt.Println("Printing map created using make ", emptyMap)

	//add some thing to emptyMap

	emptyMap["AUTH_SERVER_URI"] = "https://auth.mavenir.com"

	emptyMap["MSTORE_URI"] = "https://mstore.mavenir.com"

	fmt.Println("Printing MAP with updated values...", emptyMap)

	//resetting/clearining maps

	// Way1
	for idx, _ := range emptyMap {
		delete(emptyMap, idx)
	}

	fmt.Println("Printing map after resetting:", emptyMap)

	//Way2

	emptyMap["WRG_URL"] = "https://wrg003.mavenir.com"

	fmt.Println("Printing map after adding WRG_URL ", emptyMap)

	emptyMap = make(map[string]string)

	fmt.Println("Printing map after restting again", emptyMap)
}
