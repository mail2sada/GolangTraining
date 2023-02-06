package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps....")

	// Declating map

	var ageMap map[string]int = map[string]int{"Anil": 20, "Vishal": 21, "Bharat": 26}

	fmt.Println("Printing Age map", ageMap)

	//adding new value to the map
	ageMap["Sunil"] = 24

	fmt.Println("Printing updated age map", ageMap)

	//retrieving value

	ok, val := ageMap["Anil"]

	fmt.Println("Values returned are ", ok, val)

	ok, val = ageMap["Mark"]

	fmt.Println("Values returned are ", ok, val)

	//deleting Item from map use delete method

	delete(ageMap, "Bharat")

	fmt.Println("Printing map with delete value", ageMap)

	//finding len of map

	fmt.Println("Printing length of map", len(ageMap))

}
