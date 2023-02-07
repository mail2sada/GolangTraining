package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func division(a, b int) int {
	fmt.Println("This method will perform a/b")

	if b == 0 {
		panic("Divide by zero error")
	}

	return a / b
}

func executeDivision() {
	defer wg.Done()
	defer func() {
		res := recover()

		if res != nil {
			fmt.Println("Error occurred ...", res)
		}
	}()
	fmt.Println("result:", division(20, 2))
	fmt.Println("result:", division(40, 5))
	fmt.Println("result:", division(30, 6))
	fmt.Println("result:", division(80, 4))
	fmt.Println("result:", division(10, 0))
}

func main() {
	wg.Add(1)
	defer func() {
		res := recover()
		if res != nil {
			fmt.Println("[MAIN]Application has paniced and it is handled here, reason for panic", res)
		}
	}()

	fmt.Println("Demo: Panic and recover")

	fmt.Println("result:", division(20, 2))
	fmt.Println("result:", division(40, 5))
	fmt.Println("result:", division(30, 6))
	fmt.Println("result:", division(80, 4))
	fmt.Println("result:", division(10, 1))
	go executeDivision()
	wg.Wait()

	fmt.Println("Existing....")
}
