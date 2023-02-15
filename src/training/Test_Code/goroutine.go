package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var counter int = 0

var mutex sync.Mutex

func MyTestDisplay() {
	defer wg.Done()
	mutex.Lock()
	counter += 1
	mutex.Unlock()
	fmt.Println("From My first display...")
}

func Display() {

	defer wg.Done()
	mutex.Lock()
	counter += 1
	mutex.Unlock()
	fmt.Println("Hello, how are you?")
}

func main() {

	wg.Add(2)

	go Display()

	go MyTestDisplay()

	fmt.Println("Demo: Go routine.....")

	wg.Wait()
}
