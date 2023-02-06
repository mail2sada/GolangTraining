package main

import "fmt"

// Every thing is allowed here

/*
boolean
numeric
string,
pointer
channel  //covered later
interface types //covere later
structs –
array –
Slice
Map
Function
*/

func main() {
	bMap := map[string]bool{}
	nMap := map[string]int{}

	sMap := map[string]string{}

	pMap := map[string]*int{}

	mMap := map[string]map[string]string{}

	fMap := map[string]func(int, int) int{}

	fmt.Println(bMap, nMap, sMap, pMap, mMap, fMap)

	fmt.Printf("%v\n", bMap)

}
