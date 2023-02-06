package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func myGoroutine1() {

	fmt.Println("----------- This is from myGoroutine1....")
	defer wg.Done()
}

func myGoroutine2() {

	fmt.Println("----------- This is from myGoroutine2....")
	defer wg.Done()
}

func main() {
	wg.Add(2)
	go myGoroutine1()
	go myGoroutine2()
	wg.Wait()
	fmt.Println("Demo: Usage of Wait group..")
}
