package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo Timer...")

	t1 := time.After(5 * time.Second)

	select {

	case <-t1:
		fmt.Println("Timer expired")
	}
}
