package main

import (
	"fmt"
	"strings"
)

func main() {

	// Comparing string using Compare function
	fmt.Println(strings.Compare("mav", "Mavenir"))

	fmt.Println(strings.Compare("MavenirGoTraining",
		"MavenirGoTraining"))

	fmt.Println(strings.Compare("Mavenir", " MAV"))

	fmt.Println(strings.Compare("MaveNir", "MaveNir"))

}
