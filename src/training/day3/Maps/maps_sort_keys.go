package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Demo sorting map by keys")

	var keyValuePair = map[string]int{"Mark": 200, "Ben": 500, "Luciana": 600, "Vinod": 200, "Jacob": 50}

	var mapKeys = make([]string, 0, len(keyValuePair))

	for idx, _ := range keyValuePair {

		mapKeys = append(mapKeys, idx)
	}

	sort.Strings(mapKeys)

	for _, val := range mapKeys {

		fmt.Println(" map details for key ", val, keyValuePair[val])
	}

}
