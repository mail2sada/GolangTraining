package main

import (
	"fmt"
	"sync"
)

// main routine...

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(2)

	fmt.Println("Demo: GoRoutines....")

	go MyRoutine()

	fmt.Println("Exiting from main")
	go MyRoutine()
	waitGroup.Wait()
}

func MyRoutine() {
	defer waitGroup.Done()
	fmt.Println("In MyRoutine")
}
