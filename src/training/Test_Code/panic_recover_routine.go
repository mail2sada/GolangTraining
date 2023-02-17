package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Divide(a, b int) int {

	defer func() {
		rec := recover()

		if rec != nil {
			fmt.Println("Panic is handled...")
		}
	}()

	if b == 0 {

		fmt.Println("Panicing.... for divide by 0")

		panic("Divide by zero...")
	}

	return a / b
}

func MyTestRoutine() {

	defer wg.Done()

	fmt.Println("MyTestRoutine.....")

	var intSlice = []int{10, 20, 30, 40, 50, 0, -1, 16, 36}

	for idx, val := range intSlice {

		fmt.Println("Idx and val ", idx, val)
		Divide(idx, val)

	}

}

func main() {

	wg.Add(1)

	fmt.Println("Demo panic and recover with routine...")

	go MyTestRoutine()

	wg.Wait()

}
