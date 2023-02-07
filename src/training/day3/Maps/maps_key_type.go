package main

import "fmt"

// Allowed key types in map

/*
boolean
numeric
string,
pointer
channel
interface types
structs – if all it’s field type is comparable
array – if the type of value of array element is comparable
*/

func main() {
	fmt.Println("Demo: Allowed key_types")

	var a, b = 10, 15
	var boolMap map[bool]string = map[bool]string{}

	boolMap[true] = "a is greater than b"

	boolMap[false] = "a is not greater than b"

	fmt.Println(boolMap[a > b])

	/*
		boolean
		numeric
		string,
		pointer
		channel //revisit later
		interface types // revisit later
		structs – if all it’s field type is comparable
		array – if the type of value of array element is comparable
	*/
	type example struct {
		abc string
	}
	var bMap map[bool]int = map[bool]int{}

	var nMap map[int]string = map[int]string{}

	var sMap map[string]string = map[string]string{}

	var pMap map[*string]int = map[*string]int{}

	var structMap map[example]string = map[example]string{}

	var aMap map[[10]int]int = map[[10]int]int{}

	fmt.Println(bMap, nMap, sMap, pMap, structMap, aMap)

	

}
