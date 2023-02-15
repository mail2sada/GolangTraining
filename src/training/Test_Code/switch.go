package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Demo. Switch")

	switch t := time.Now(); {

	case t.Day() <= 8:

		fmt.Println("First week")
	case t.Day() <= 14:

		fmt.Println("Second week")

	}

}
