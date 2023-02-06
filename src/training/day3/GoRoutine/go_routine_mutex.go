package main

import (
	"fmt"
	"sync"
)

var (
	counter int32
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func increment(msg string) {
	defer wg.Done()

	for i := 0; i < 3; i++ {

		mutex.Lock()
		{
			fmt.Println(msg)
			counter++
		}
		mutex.Unlock()
	}
}

func main() {

	fmt.Println("Demo: Sync Mutex...")
	wg.Add(3)

	go increment("Welcome ")

	go increment("Go lang Training")

	go increment("Hello Mavenir...")

	wg.Wait()

	fmt.Println("Value of counter after increment ", counter)

}
