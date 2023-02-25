package main

import "fmt"

func main() {

	fmt.Println("Demo: Maps....")

	var myMap map[string]string = map[string]string{"Suresh": "Hello", "Yogesh": "Hi"}

	fmt.Println("Length of map is ", len(myMap))

	myMap["Vishal"] = "Android team"
	myMap["Dilip"] = "iOS team"
	myMap["Sada"] = ""
	a := myMap["Dilip"]

	x, avl := myMap["Sada"]

	fmt.Println(avl)

	fmt.Println(x)

	fmt.Println(a)

	for key, val := range myMap {

		fmt.Println("Key is ", key, "Value is ", val)
	}

	var heapMap map[string]int = make(map[string]int)

	delete(myMap, "Sada")

	for key, val := range myMap {

		fmt.Println("Key is ", key, "Value is ", val)
	}

	for key, _ := range myMap {
		delete(myMap, key)
	}

	myMap = make(map[string]string)

	fmt.Println(heapMap)

}
