package main

import (
	"fmt"
	"time"
)

// important address scope of timer in training,,,,

func main() {

	fmt.Println("Demo: timer.AfterFunc")

	var a, b int = 10, 20

	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("Timer expired...", a, b)
	})

	defer timer.Stop()

	time.Sleep(10 * time.Second)
}
