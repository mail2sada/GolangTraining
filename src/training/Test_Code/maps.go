package main

import "fmt"

func main() {

	fmt.Println("Demo: Maps....")

	var firstMap = map[string]string{}

	firstMap["WrgUrl"] = "https://mavenir.wrg.com:443/resourceURL"

	firstMap["MstoreUrl"] = "https://mavenir.mstore.com:443/resourceURL"

	fmt.Println(firstMap["WrgUrl"])

	fmt.Println(firstMap["MstoreUrl"])

	fmt.Println("length of map is ", len(firstMap))

	for key, val := range firstMap {

		fmt.Println("Key:", key, "value:", val)
	}

	employee := make(map[string]int)

	fmt.Println("len of map emp", len(employee))

	employee["Anil"] = 10

	employee["Vijay"] = 20

	val, err := employee["Bhanu"]

	fmt.Println("Value os bhanu is", val, err)

	employee["Vijay"] = 500

	for key, val := range employee {

		fmt.Println("Key:", key, "value:", val)
	}

	//delete item

	delete(employee, "Vijay")

	for key, val := range employee {

		fmt.Println("Key:", key, "value:", val)
	}

}
